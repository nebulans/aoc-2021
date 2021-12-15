import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "1163751742",
    "1381373672",
    "2136511328",
    "3694931569",
    "7463417111",
    "1319128137",
    "1359912421",
    "3125421639",
    "1293138521",
    "2311944581",
]
ACTUAL_INPUT = 'input/day_15.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(15, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '40'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '462'


@pytest.mark.parametrize('solution', [
    GoSolution(15, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '315'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '2846'
