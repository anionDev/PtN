from ScriptCollection.TFCPS.TFCPS_CodeUnitSpecific_Base import TFCPS_CodeUnitSpecific_Base,TFCPS_CodeUnitSpecific_Base_CLI
from ScriptCollection.TFCPS.Go.TFCPS_CodeUnitSpecific_Go import TFCPS_CodeUnitSpecific_Go_Functions , TFCPS_CodeUnitSpecific_Go_CLI

def run_testcases():
    t :TFCPS_CodeUnitSpecific_Go_Functions= TFCPS_CodeUnitSpecific_Go_CLI().parse(__file__)
    t.run_testcases()


if __name__ == "__main__":
    run_testcases()
