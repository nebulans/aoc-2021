import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "6,10",
    "0,14",
    "9,10",
    "0,3",
    "10,4",
    "4,11",
    "6,0",
    "6,12",
    "4,1",
    "0,13",
    "10,12",
    "3,4",
    "3,0",
    "8,4",
    "1,10",
    "2,14",
    "8,10",
    "9,0",
    "",
    "fold along y=7",
    "fold along x=5",
]
ACTUAL_INPUT = 'input/day_13.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(13, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '17'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '704'


@pytest.mark.parametrize('solution', [
    GoSolution(13, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '16'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '103'
