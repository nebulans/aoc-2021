import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "0,9 -> 5,9",
    "8,0 -> 0,8",
    "9,4 -> 3,4",
    "2,2 -> 2,1",
    "7,0 -> 7,4",
    "6,4 -> 2,0",
    "0,9 -> 2,9",
    "3,4 -> 1,4",
    "0,0 -> 8,8",
    "5,5 -> 8,2",
]
ACTUAL_INPUT = 'input/day_05.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(5, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '5'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '6283'


@pytest.mark.parametrize('solution', [
    GoSolution(5, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '12'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '18864'
