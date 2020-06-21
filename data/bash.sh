function _sup() {
    PS1=$(sup ${PIPESTATUS[@]})
}

PROMPT_COMMAND=_sup
