import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "5483143223",
    "2745854711",
    "5264556173",
    "6141336146",
    "6357385478",
    "4167524645",
    "2176841721",
    "6882881134",
    "4846848554",
    "5283751526"
]
ACTUAL_INPUT = 'input/day_11.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(11, 1),
    GoSolution("11m", 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '1656'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1601'


@pytest.mark.parametrize('solution', [
    GoSolution(11, 2),
    GoSolution("11m", 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '195'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '368'
