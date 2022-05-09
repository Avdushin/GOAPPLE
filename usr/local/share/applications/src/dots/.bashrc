fish

#noisetorch

if [ -d "$HOME/.local/bin" ] ; then
    PATH="$HOME/.local/bin:$PATH"
fi

echo " 
#fish_config
if status is-interactive
    # Commands to run in interactive sessions can go here
set -g fish_greeting
end " > ~/.config/fish/config.fish

echo " 
function gt
        command cat ~/git_token
end
    
function subl
    command .appz/sublime_text/sublime_text
end" >> ~/.config/fish/functions/fish_prompt.fish