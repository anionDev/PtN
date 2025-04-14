import sys
from pathlib import Path
from ScriptCollection.TasksForCommonProjectStructure import TasksForCommonProjectStructure


def start_dockerfile_example():
    TasksForCommonProjectStructure().start_dockerfile_example(str(Path(__file__).absolute()), 3, True, True, sys.argv)


if __name__ == "__main__":
    start_dockerfile_example()
