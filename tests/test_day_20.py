import pytest

from . import GoSolution

OFFICIAL_EXAMPLE = 'input/day_20_example.txt'
ACTUAL_INPUT = 'input/day_20.txt'


@pytest.mark.parametrize('solution', [
    GoSolution(20, 1),
], ids=lambda s: s.name)
class TestPart1(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '35'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '5419'


@pytest.mark.parametrize('solution', [
    GoSolution(20, 2),
], ids=lambda s: s.name)
class TestPart2(object):

    def test_official_example(self, solution):
        assert solution.run_file(OFFICIAL_EXAMPLE) == '3351'

    def test_actual_input(self, solution):
        assert solution.run_file(ACTUAL_INPUT) == '17325'
