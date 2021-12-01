import abc
import subprocess


class Solution(abc.ABC):
    entrypoint = []

    def __init__(self, day, part):
        self.day = day
        self.part = part

    def get_command(self, **kwargs):
        return self.entrypoint + [str(self.day), str(self.part)] + [f'--{k.replace("_", "-")}={v}' for k, v in kwargs.items()]

    def run(self, input, **kwargs):
        command = self.get_command(**kwargs)
        result = subprocess.run(command, stdout=subprocess.PIPE, input=input, timeout=120, encoding='utf-8', check=True)
        return result.stdout.strip()

    def run_file(self, path, **kwargs):
        with open(path) as fd:
            return self.run(fd.read(), **kwargs)

    def run_lines(self, lines, **kwargs):
        return self.run('\n'.join(str(l) for l in lines), **kwargs)


class PythonSolution(Solution):
    name = 'Python'
    entrypoint = ['python3', 'solutions/python/__init__.py']


class JavaSolution(Solution):
    name = 'Java'
    entrypoint = ['java', '-cp', 'solutions/java/target/java-1.0-SNAPSHOT.jar', 'xyz.timdavies.aoc2020.App']


class GoSolution(Solution):
    name = 'Go'
    entrypoint = ['go/aog']
