import curses

def menu(table, title = "Select an option:"):
    def inner(stdscr):
        curses.curs_set(0)
        options = list(table) + ["Quit"]
        current_row = 0
        curses.init_pair(1, curses.COLOR_BLACK, curses.COLOR_WHITE)
        while True:
            stdscr.clear()
            stdscr.addstr(0, 0, title + "\n", curses.A_BOLD)
            for i, option in enumerate(options):
                if i == current_row:
                    stdscr.attron(curses.color_pair(1))
                    stdscr.addstr(i + 1, 0, f"> {option}")
                    stdscr.attroff(curses.color_pair(1))
                else:
                    stdscr.addstr(i + 1, 0, f"  {option}")
            key = stdscr.getch()
            if key == curses.KEY_UP and current_row > 0:
                current_row -= 1
            elif key == curses.KEY_DOWN and current_row < len(options) - 1:
                current_row += 1
            elif key in [10, 13]:  # Enter
                if options[current_row] == "Quit":
                    return None
                return options[current_row]
            stdscr.refresh()
    return curses.wrapper(inner)