#!/usr/bin/osascript

set currentDirectory to "$HOME/workspace/otel-spike/"

# Open iTerm2 and split it into two panes
tell application "iTerm"
    activate
    
    # Create window
    set newWindow to (create window with default profile)
    
    # Split pane
    tell current session of newWindow
        split horizontally with default profile
    end tell

    # Set working directory for both sessions and execute the commands
    tell first session of current tab of newWindow
        write text "cd " & currentDirectory & " && bash -c '" & currentDirectory & "/infrastructure/kafka/debug.sh send'"
    end tell

    tell second session of current tab of newWindow
        write text "cd " & currentDirectory & " && bash -c '" & currentDirectory & "/infrastructure/kafka/debug.sh consume'"
    end tell
end tell
