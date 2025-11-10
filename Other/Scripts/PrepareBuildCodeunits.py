from pathlib import Path
from ScriptCollection.GeneralUtilities import GeneralUtilities
from ScriptCollection.TFCPS.TFCPS_Generic import TFCPS_Generic_Functions, TFCPS_Generic_CLI
    
def prepare_build_codeunits():
    t :TFCPS_Generic_Functions= TFCPS_Generic_CLI().parse(__file__)
    repository_folder = GeneralUtilities.resolve_relative_path( "../../..", str(Path(__file__).absolute()))
    t.tfcps_Tools_General.ensure_certificate_authority_for_development_purposes_is_generated(repository_folder)
    t.tfcps_Tools_General.generate_certificate_for_development_purposes_for_product(repository_folder)
    t.tfcps_Tools_General.generate_tasksfile_from_workspace_file(repository_folder)
    t.tfcps_Tools_General.generate_codeunits_overview_diagram(repository_folder)
    t.tfcps_Tools_General.generate_svg_files_from_plantuml_files_for_repository(repository_folder,t.use_cache())


if __name__ == "__main__":
    prepare_build_codeunits()
