import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [
    "Player 1 starting position: 4",
    "Player 2 starting position: 8",
]
ACTUAL_INPUT = 'input/day_21.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(21, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '739785'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '908091'


@pytest.mark.parametrize('solution', [
    GoSolution(21, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '444356092776315'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '190897246590017'
