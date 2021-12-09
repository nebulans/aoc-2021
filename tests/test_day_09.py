import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "2199943210",
    "3987894921",
    "9856789892",
    "8767896789",
    "9899965678",
]
ACTUAL_INPUT = 'input/day_09.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(9, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '15'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '518'


@pytest.mark.parametrize('solution', [
    GoSolution(9, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '1134'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '949905'
