def better_input(prompt: str, default_value: str | None = None, allow_empty: bool = False) -> str | None:
    """
    Ask the user until they provide a valid input

    Arguments: 
        prompt: str. The message/question you want to display
        default_value: str or None (default None). A default value if the user presses Enter without typing anything
        allow_empty: bool (default False). Whether to allow empty input
    
    Returns:
        str: if the user types something
        None: if the user presses Enter and allow_empty is True or a default_value is None
    """
    while True:
        user_input = input(prompt).strip()
        
        if user_input != "" and default_value is None and allow_empty is False:
            return user_input

        if not user_input and default_value is not None:
            return default_value
        
        if allow_empty and not user_input:
            return None