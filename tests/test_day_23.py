import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "#############",
    "#...........#",
    "###B#C#B#D###",
    "  #A#D#C#A#",
    "  #########",
]
OFFICIAL_EXAMPLE_FULL = [
    "#############",
    "#...........#",
    "###B#C#B#D###",
    "  #D#C#B#A#",
    "  #D#B#A#C#",
    "  #A#D#C#A#",
    "  #########",
]
ACTUAL_INPUT = 'input/day_23.txt'
ACTUAL_INPUT_FULL = 'input/day_23_full.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(23, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '12521'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '10526'


@pytest.mark.parametrize('solution', [
    GoSolution(23, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE_FULL) == '44169'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT_FULL) == '41284'
