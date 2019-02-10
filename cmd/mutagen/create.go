package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	fs "github.com/havoc-io/mutagen/pkg/filesystem"
	promptpkg "github.com/havoc-io/mutagen/pkg/prompt"
	sessionsvcpkg "github.com/havoc-io/mutagen/pkg/service/session"
	sessionpkg "github.com/havoc-io/mutagen/pkg/session"
	"github.com/havoc-io/mutagen/pkg/sync"
	"github.com/havoc-io/mutagen/pkg/url"
)

func createMain(command *cobra.Command, arguments []string) error {
	// Validate, extract, and parse URLs.
	if len(arguments) != 2 {
		return errors.New("invalid number of endpoint URLs provided")
	}
	alpha, err := url.Parse(arguments[0], true)
	if err != nil {
		return errors.Wrap(err, "unable to parse alpha URL")
	}
	beta, err := url.Parse(arguments[1], false)
	if err != nil {
		return errors.Wrap(err, "unable to parse beta URL")
	}

	// If either URL is a local path, make sure it's normalized.
	if alpha.Protocol == url.Protocol_Local {
		if alphaPath, err := fs.Normalize(alpha.Path); err != nil {
			return errors.Wrap(err, "unable to normalize alpha path")
		} else {
			alpha.Path = alphaPath
		}
	}
	if beta.Protocol == url.Protocol_Local {
		if betaPath, err := fs.Normalize(beta.Path); err != nil {
			return errors.Wrap(err, "unable to normalize beta path")
		} else {
			beta.Path = betaPath
		}
	}

	// Validate and convert the synchronization mode specification.
	var synchronizationMode sync.SynchronizationMode
	if createConfiguration.synchronizationMode != "" {
		if err := synchronizationMode.UnmarshalText([]byte(createConfiguration.synchronizationMode)); err != nil {
			return errors.Wrap(err, "unable to parse synchronization mode")
		}
	}

	// Validate and convert the symbolic link mode specification.
	var symbolicLinkMode sync.SymlinkMode
	if createConfiguration.symbolicLinkMode != "" {
		if err := symbolicLinkMode.UnmarshalText([]byte(createConfiguration.symbolicLinkMode)); err != nil {
			return errors.Wrap(err, "unable to parse symbolic link mode")
		}
	}

	// Validate and convert the watch mode specification.
	var watchMode fs.WatchMode
	if createConfiguration.watchMode != "" {
		if err := watchMode.UnmarshalText([]byte(createConfiguration.watchMode)); err != nil {
			return errors.Wrap(err, "unable to parse watch mode")
		}
	}

	// NOTE: There's no need to validate the watch polling interval - any uint32
	// value is valid.

	// NOTE: We don't need to validate ignores here, that will happen on the
	// session service, so we'll save ourselves the time.

	// Validate and convert the VCS ignore mode specification.
	var ignoreVCSMode sync.IgnoreVCSMode
	if createConfiguration.ignoreVCS && createConfiguration.noIgnoreVCS {
		return errors.New("conflicting VCS ignore behavior specified")
	} else if createConfiguration.ignoreVCS {
		ignoreVCSMode = sync.IgnoreVCSMode_IgnoreVCS
	} else if createConfiguration.noIgnoreVCS {
		ignoreVCSMode = sync.IgnoreVCSMode_PropagateVCS
	}

	// Validate and convert default file mode specifications.
	var defaultFileMode, defaultFileModeAlpha, defaultFileModeBeta uint32
	if createConfiguration.defaultFileMode != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultFileMode, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default file mode")
		} else if err = sync.EnsureDefaultFileModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default file mode")
		} else {
			defaultFileMode = uint32(m)
		}
	}
	if createConfiguration.defaultFileModeAlpha != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultFileModeAlpha, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default file mode for alpha")
		} else if err = sync.EnsureDefaultFileModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default file mode for alpha")
		} else {
			defaultFileModeAlpha = uint32(m)
		}
	}
	if createConfiguration.defaultFileModeBeta != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultFileModeBeta, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default file mode for beta")
		} else if err = sync.EnsureDefaultFileModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default file mode for beta")
		} else {
			defaultFileModeBeta = uint32(m)
		}
	}

	// Validate and convert default directory mode specifications.
	var defaultDirectoryMode, defaultDirectoryModeAlpha, defaultDirectoryModeBeta uint32
	if createConfiguration.defaultDirectoryMode != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultDirectoryMode, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default directory mode")
		} else if err = sync.EnsureDefaultDirectoryModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default directory mode")
		} else {
			defaultDirectoryMode = uint32(m)
		}
	}
	if createConfiguration.defaultDirectoryModeAlpha != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultDirectoryModeAlpha, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default directory mode for alpha")
		} else if err = sync.EnsureDefaultDirectoryModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default directory mode for alpha")
		} else {
			defaultDirectoryModeAlpha = uint32(m)
		}
	}
	if createConfiguration.defaultDirectoryModeBeta != "" {
		if m, err := fs.ParseMode(createConfiguration.defaultDirectoryModeBeta, fs.ModePermissionsMask); err != nil {
			return errors.Wrap(err, "unable to parse default directory mode for beta")
		} else if err = sync.EnsureDefaultDirectoryModeValid(m); err != nil {
			return errors.Wrap(err, "invalid default directory mode for beta")
		} else {
			defaultDirectoryModeBeta = uint32(m)
		}
	}

	// Validate default file owner user specifications.
	if createConfiguration.defaultUser != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultUser); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid user ownership specification")
		}
	}
	if createConfiguration.defaultUserAlpha != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultUserAlpha); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid user ownership specification for alpha")
		}
	}
	if createConfiguration.defaultUserBeta != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultUserBeta); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid user ownership specification for beta")
		}
	}

	// Validate default file owner group specifications.
	if createConfiguration.defaultGroup != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultGroup); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid group ownership specification")
		}
	}
	if createConfiguration.defaultGroupAlpha != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultGroupAlpha); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid group ownership specification for alpha")
		}
	}
	if createConfiguration.defaultGroupBeta != "" {
		if kind, _ := fs.ParseOwnershipIdentifier(createConfiguration.defaultGroupBeta); kind == fs.OwnershipIdentifierKindInvalid {
			return errors.New("invalid group ownership specification for beta")
		}
	}

	// Connect to the daemon and defer closure of the connection.
	daemonConnection, err := createDaemonClientConnection()
	if err != nil {
		return errors.Wrap(err, "unable to connect to daemon")
	}
	defer daemonConnection.Close()

	// Create a session service client.
	sessionService := sessionsvcpkg.NewSessionsClient(daemonConnection)

	// Invoke the session create method. The stream will close when the
	// associated context is cancelled.
	createContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := sessionService.Create(createContext)
	if err != nil {
		return errors.Wrap(peelAwayRPCErrorLayer(err), "unable to invoke create")
	}

	// Send the initial request.
	request := &sessionsvcpkg.CreateRequest{
		Alpha: alpha,
		Beta:  beta,
		Configuration: &sessionpkg.Configuration{
			SynchronizationMode:                 synchronizationMode,
			SymlinkMode:                         symbolicLinkMode,
			WatchMode:                           watchMode,
			WatchPollingInterval:                createConfiguration.watchPollingInterval,
			Ignores:                             createConfiguration.ignores,
			IgnoreVCSMode:                       ignoreVCSMode,
			PermissionDefaultFileMode:           defaultFileMode,
			PermissionDefaultFileModeAlpha:      defaultFileModeAlpha,
			PermissionDefaultFileModeBeta:       defaultFileModeBeta,
			PermissionDefaultDirectoryMode:      defaultDirectoryMode,
			PermissionDefaultDirectoryModeAlpha: defaultDirectoryModeAlpha,
			PermissionDefaultDirectoryModeBeta:  defaultDirectoryModeBeta,
			PermissionDefaultUser:               createConfiguration.defaultUser,
			PermissionDefaultUserAlpha:          createConfiguration.defaultUserAlpha,
			PermissionDefaultUserBeta:           createConfiguration.defaultUserBeta,
			PermissionDefaultGroup:              createConfiguration.defaultGroup,
			PermissionDefaultGroupAlpha:         createConfiguration.defaultGroupAlpha,
			PermissionDefaultGroupBeta:          createConfiguration.defaultGroupBeta,
		},
	}
	if err := stream.Send(request); err != nil {
		return errors.Wrap(peelAwayRPCErrorLayer(err), "unable to send create request")
	}

	// Create a status line printer and defer a break.
	statusLinePrinter := &cmd.StatusLinePrinter{}
	defer statusLinePrinter.BreakIfNonEmpty()

	// Receive and process responses until we're done.
	for {
		if response, err := stream.Recv(); err != nil {
			return errors.Wrap(peelAwayRPCErrorLayer(err), "create failed")
		} else if err = response.EnsureValid(); err != nil {
			return errors.Wrap(err, "invalid create response received")
		} else if response.Session != "" {
			statusLinePrinter.Print(fmt.Sprintf("Created session %s", response.Session))
			return nil
		} else if response.Message != "" {
			statusLinePrinter.Print(response.Message)
			if err := stream.Send(&sessionsvcpkg.CreateRequest{}); err != nil {
				return errors.Wrap(peelAwayRPCErrorLayer(err), "unable to send message response")
			}
		} else if response.Prompt != "" {
			statusLinePrinter.BreakIfNonEmpty()
			if response, err := promptpkg.PromptCommandLine(response.Prompt); err != nil {
				return errors.Wrap(err, "unable to perform prompting")
			} else if err = stream.Send(&sessionsvcpkg.CreateRequest{Response: response}); err != nil {
				return errors.Wrap(peelAwayRPCErrorLayer(err), "unable to send prompt response")
			}
		}
	}
}

var createCommand = &cobra.Command{
	Use:   "create <alpha> <beta>",
	Short: "Creates and starts a new synchronization session",
	Run:   cmd.Mainify(createMain),
}

var createConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
	// synchronizationMode specifies the synchronization mode for the session.
	synchronizationMode string
	// symbolicLinkMode specifies the symbolic link handling mode to use for
	// the session.
	symbolicLinkMode string
	// watchMode specifies the filesystem watching mode to use for the session.
	watchMode string
	// watchPollingInterval specifies the polling interval to use if using
	// poll-based or hybrid watching.
	watchPollingInterval uint32
	// ignores is the list of ignore specifications for the session.
	ignores []string
	// ignoreVCS specifies whether or not to enable VCS ignores for the session.
	ignoreVCS bool
	// noIgnoreVCS specifies whether or not to disable VCS ignores for the
	// session.
	noIgnoreVCS bool
	// defaultFileMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode, with endpoint-specific
	// specifications taking priority.
	defaultFileMode string
	// defaultFileModeAlpha specifies the default permission mode to use for new
	// files on alpha in "portable" permission propagation mode, taking priority
	// over defaultFileMode on alpha if specified.
	defaultFileModeAlpha string
	// defaultFileModeBeta specifies the default permission mode to use for new
	// files on beta in "portable" permission propagation mode, taking priority
	// over defaultFileMode on beta if specified.
	defaultFileModeBeta string
	// defaultDirectoryMode specifies the default permission mode to use for new
	// directories in "portable" permission propagation mode, with endpoint-
	// specific specifications taking priority.
	defaultDirectoryMode string
	// defaultDirectoryModeAlpha specifies the default permission mode to use
	// for new directories on alpha in "portable" permission propagation mode,
	// taking priority over defaultDirectoryMode on alpha if specified.
	defaultDirectoryModeAlpha string
	// defaultDirectoryModeBeta specifies the default permission mode to use for
	// new directories on beta in "portable" permission propagation mode, taking
	// priority over defaultDirectoryMode on beta if specified.
	defaultDirectoryModeBeta string
	// defaultUser specifies the default user identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode, with endpoint-specific specifications taking priority.
	defaultUser string
	// defaultUserAlpha specifies the default user identifier to use when
	// setting ownership of new files and directories on alpha in "portable"
	// permission propagation mode, taking priority over defaultUser on alpha if
	// specified.
	defaultUserAlpha string
	// defaultUserBeta specifies the default user identifier to use when setting
	// ownership of new files and directories on beta in "portable" permission
	// propagation mode, taking priority over defaultUser on beta if specified.
	defaultUserBeta string
	// defaultGroup specifies the default group identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode, with endpoint-specific specifications taking priority.
	defaultGroup string
	// defaultGroupAlpha specifies the default group identifier to use when
	// setting ownership of new files and directories on alpha in "portable"
	// permission propagation mode, taking priority over defaultGroup on alpha
	// if specified.
	defaultGroupAlpha string
	// defaultGroupBeta specifies the default group identifier to use when
	// setting ownership of new files and directories on beta in "portable"
	// permission propagation mode, taking priority over defaultGroup on beta if
	// specified.
	defaultGroupBeta string
}

func init() {
	// Grab a handle for the command line flags.
	flags := createCommand.Flags()

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&createConfiguration.help, "help", "h", false, "Show help information")

	// Wire up synchronization flags.
	flags.StringVarP(&createConfiguration.synchronizationMode, "sync-mode", "m", "", "Specify synchronization mode (symmetric|source-wins|mirror-safe|mirror-exact)")

	// Wire up symbolic link flags.
	flags.StringVar(&createConfiguration.symbolicLinkMode, "symlink-mode", "", "Specify symlink mode (ignore|portable|posix-raw)")

	// Wire up watch flags.
	flags.StringVar(&createConfiguration.watchMode, "watch-mode", "", "Specify watch mode (portable|force-poll)")
	flags.Uint32Var(&createConfiguration.watchPollingInterval, "watch-polling-interval", 0, "Specify watch polling interval in seconds")

	// Wire up ignore flags.
	flags.StringSliceVarP(&createConfiguration.ignores, "ignore", "i", nil, "Specify ignore paths")
	flags.BoolVar(&createConfiguration.ignoreVCS, "ignore-vcs", false, "Ignore VCS directories")
	flags.BoolVar(&createConfiguration.noIgnoreVCS, "no-ignore-vcs", false, "Propagate VCS directories")

	// Wire up permission flags.
	flags.StringVar(&createConfiguration.defaultFileMode, "default-file-mode", "", "Specify default file permission mode")
	flags.StringVar(&createConfiguration.defaultFileModeAlpha, "default-file-mode-alpha", "", "Specify default file permission mode for alpha")
	flags.StringVar(&createConfiguration.defaultFileModeBeta, "default-file-mode-beta", "", "Specify default file permission mode for beta")
	flags.StringVar(&createConfiguration.defaultDirectoryMode, "default-directory-mode", "", "Specify default directory permission mode")
	flags.StringVar(&createConfiguration.defaultDirectoryModeAlpha, "default-directory-mode-alpha", "", "Specify default directory permission mode for alpha")
	flags.StringVar(&createConfiguration.defaultDirectoryModeBeta, "default-directory-mode-beta", "", "Specify default directory permission mode for beta")
	flags.StringVar(&createConfiguration.defaultUser, "default-owner-user", "", "Specify default file owner user")
	flags.StringVar(&createConfiguration.defaultUserAlpha, "default-owner-user-alpha", "", "Specify default file owner user for alpha")
	flags.StringVar(&createConfiguration.defaultUserBeta, "default-owner-user-beta", "", "Specify default file owner user for beta")
	flags.StringVar(&createConfiguration.defaultGroup, "default-owner-group", "", "Specify default file owner group")
	flags.StringVar(&createConfiguration.defaultGroupAlpha, "default-owner-group-alpha", "", "Specify default file owner group for alpha")
	flags.StringVar(&createConfiguration.defaultGroupBeta, "default-owner-group-beta", "", "Specify default file owner group for beta")
}
