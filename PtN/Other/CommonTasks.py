import sys
import os
from pathlib import Path
from ScriptCollection.ScriptCollectionCore import ScriptCollectionCore
from ScriptCollection.GeneralUtilities import GeneralUtilities
from ScriptCollection.TasksForCommonProjectStructure import TasksForCommonProjectStructure


def common_tasks():
    cmd_args = sys.argv
    t = TasksForCommonProjectStructure()
    sc = ScriptCollectionCore()
    build_environment = t.get_targetenvironmenttype_from_commandline_arguments(cmd_args, "QualityCheck")
    verbosity = t.get_verbosity_from_commandline_arguments(cmd_args, 1)
    file = str(Path(__file__).absolute())
    codeunitname = os.path.basename(GeneralUtilities.resolve_relative_path("..", os.path.dirname(file)))
    codeunit_folder = GeneralUtilities.resolve_relative_path("..", os.path.dirname(file))
    repository_folder = GeneralUtilities.resolve_relative_path("..", codeunit_folder)
    codeunit_version = sc.get_semver_version_from_gitversion(repository_folder)  # Should always be the same as the project-version
    folder_of_current_file = os.path.dirname(file)
    sc.replace_version_in_dockerfile_file(GeneralUtilities.resolve_relative_path(f"../{codeunitname}/Dockerfile", folder_of_current_file), codeunit_version)
    additional_arguments_file = t.get_additionalargumentsfile_from_commandline_arguments(cmd_args, None)
    t.standardized_tasks_do_common_tasks(file, codeunit_version, verbosity, build_environment, True, additional_arguments_file, False, cmd_args)
    t.standardized_tasks_update_version_in_docker_examples(file, codeunit_version)


if __name__ == "__main__":
    common_tasks()
