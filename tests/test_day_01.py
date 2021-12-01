import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
ACTUAL_INPUT = 'input/day_01.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(1, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '7'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1655'


@pytest.mark.parametrize('solution', [
    GoSolution(1, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_lines(OFFICIAL_EXAMPLE) == '5'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '1683'
