#!/usr/bin/zsh

function _sup() {
    PROMPT=$(sup -sh zsh ${PIPESTATUS[@]})
}

function install_sup() {
  for s in "${precmd_functions[@]}"; do
    if [ "$s" = "_sup" ]; then
      return
    fi
  done
  precmd_functions+=(_sup)
}

install_sup
