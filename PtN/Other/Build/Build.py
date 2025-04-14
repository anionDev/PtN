import sys
from pathlib import Path
from ScriptCollection.TasksForCommonProjectStructure import TasksForCommonProjectStructure


def build():
    TasksForCommonProjectStructure().standardized_tasks_build_for_docker_project(str(Path(__file__).absolute()), "QualityCheck", 1, sys.argv)


if __name__ == "__main__":
    build()
