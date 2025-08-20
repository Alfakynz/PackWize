def better_input(prompt, default=None, allow_empty=False):
    while True:
        user_input = input(prompt).strip()
        
        if user_input != "" and default is None and allow_empty is False:
            return user_input

        if not user_input and default is not None:
            return default
        
        if allow_empty and not user_input:
            return None