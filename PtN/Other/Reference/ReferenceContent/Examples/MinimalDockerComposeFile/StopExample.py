import sys
from pathlib import Path
from ScriptCollection.TasksForCommonProjectStructure import TasksForCommonProjectStructure


def stop_dockerfile_example():
    TasksForCommonProjectStructure().stop_dockerfile_example(str(Path(__file__).absolute()), 3, True, True, sys.argv)


if __name__ == "__main__":
    stop_dockerfile_example()
