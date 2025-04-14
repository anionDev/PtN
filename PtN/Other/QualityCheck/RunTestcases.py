import os
from pathlib import Path
from ScriptCollection.GeneralUtilities import GeneralUtilities
from ScriptCollection.ScriptCollectionCore import ScriptCollectionCore


def run_testcases():
    current_file: str = str(Path(__file__).resolve())
    codeunit_folder = GeneralUtilities.resolve_relative_path("../../..", current_file)
    codeunit_name: str = os.path.basename(codeunit_folder)
    test_coverage_folder = os.path.join(codeunit_folder, "Other", "Artifacts", "TestCoverage").replace("\\", "/")
    GeneralUtilities.ensure_directory_exists(test_coverage_folder)
    src_folder = GeneralUtilities.resolve_relative_path(codeunit_name, codeunit_folder)
    sc: ScriptCollectionCore = ScriptCollectionCore()
    sc.run_program_argsasarray("go", ["install", "github.com/t-yuki/gocover-cobertura@latest"], src_folder)
    sc.run_program_argsasarray("go", ["test", "-coverprofile=coverage.out", "./..."], src_folder)
    sc.run_program_argsasarray("sh", ["-c", f"gocover-cobertura < coverage.out > {test_coverage_folder}/coverage.xml"], src_folder)
    # TODO do common follow-up-tasks for testcoverage-file (verification, generate html-documentation, etc.)


if __name__ == "__main__":
    run_testcases()
