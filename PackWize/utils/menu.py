import curses

def menu(table: list[str], title: str = "Select an option:"):
    """
    Menu for the TUI.

    Arguments:
        table: list[str]. All available options
        title: str (default: 'Select an option:'). The menu title
    
    Returns:
        str: the option selected
        None: if an error occurs or no option is selected
    """
    def inner(stdscr):
        curses.curs_set(0)
        options = list(table) + ["Quit"]
        current_row = 0
        curses.init_pair(1, curses.COLOR_BLACK, curses.COLOR_WHITE)

        while True:
            stdscr.clear()
            h, w = stdscr.getmaxyx()

            max_visible = h - 2
            if max_visible <= 0:
                msg = "Terminal too small! Resize window."
                stdscr.addstr(0, 0, msg[:w-1], curses.A_BOLD)
                stdscr.refresh()
                key = stdscr.getch()
                if key in [ord("q"), 27]:
                    return None
                continue

            # Calculate window
            start = max(0, current_row - max_visible + 1) if current_row >= max_visible else 0
            end = min(len(options), start + max_visible)

            # Title
            stdscr.addstr(0, 0, title[:w-2], curses.A_BOLD)

            # Options
            for idx, option in enumerate(options[start:end], start=start):
                y = (idx - start) + 1
                text = f"{idx+1}. {option}"  # Number
                if idx == current_row:
                    stdscr.addstr(y, 0, text[:w-2], curses.color_pair(1))
                else:
                    stdscr.addstr(y, 0, text[:w-2])

            # Scroll bar
            if len(options) > max_visible:
                bar_height = max(1, max_visible * max_visible // len(options))
                bar_pos = int((current_row / (len(options) - 1)) * (max_visible - bar_height))
                for i in range(max_visible):
                    char = "█" if bar_pos <= i < bar_pos + bar_height else "│"
                    stdscr.addstr(i + 1, w - 1, char)

            stdscr.refresh()
            key = stdscr.getch()

            # Navigation
            if key == curses.KEY_UP and current_row > 0:
                current_row -= 1
            elif key == curses.KEY_DOWN and current_row < len(options) - 1:
                current_row += 1
            elif key in [10, 13]:  # Enter
                if options[current_row] == "Quit":
                    return None
                return options[current_row]
            elif ord("0") <= key <= ord("9"):  
                # Naviagation using number keys
                num = key - ord("0")
                if num > 0 and num <= len(options):
                    current_row = num - 1

    return curses.wrapper(inner)
