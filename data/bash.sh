function _sup() {
    PS1=$(sup ${PIPESTATUS[@]})
}

if [[ -z ${VTE_VERSION} ]]; then
    PROMPT_COMMAND="_sup"
else
    PROMPT_COMMAND="_sup; __vte_prompt_command"
fi
