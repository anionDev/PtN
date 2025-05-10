import requests
from packaging import version
import json
import re
from pathlib import Path
from ScriptCollection.GeneralUtilities import GeneralUtilities


def get_latest_version(version_strings: list[str]) -> str:
    parsed = [version.parse(v) for v in version_strings]
    result = max(parsed)
    return str(result)


def get_latest_go_version_in_docker_alpine_image() -> str:
    response = requests.get("https://hub.docker.com/v2/repositories/library/golang/tags?name=alpine&ordering=last_updated&page=1&page_size=250", timeout=20, headers={'Cache-Control': 'no-cache'})
    if response.status_code != 200:
        raise ValueError(f"Failed to fetch data from Docker Hub: {response.status_code}")
    response_text = response.text
    data = json.loads(response_text)
    tags: list[str] = [tag["name"] for tag in data["results"] if re.match(r'^\d+\.\d+\.\d+\-alpine$', tag["name"])]
    versions = [tag.split("-")[0] for tag in tags]
    latest_version = get_latest_version(versions)
    return latest_version


def update_dependencies():
    latest_go_version = get_latest_go_version_in_docker_alpine_image()
    current_file_path = str(Path(__file__))
    codeunit_folder: str = GeneralUtilities.resolve_relative_path("../..", current_file_path)
    docker_file: str = GeneralUtilities.resolve_relative_path("PtN/Dockerfile", codeunit_folder)
    GeneralUtilities.replace_regex_in_file(docker_file, "FROM golang:\\d+\\.\\d+\\.\\d+\\-alpine AS builder", f"FROM golang:{latest_go_version}-alpine AS builder")
    go_mod_file: str = GeneralUtilities.resolve_relative_path("PtN/go.mod", codeunit_folder)
    GeneralUtilities.replace_regex_in_file(go_mod_file, "go \\d+\\.\\d+\\.\\d+", f"go {latest_go_version}")
    readme_file: str = GeneralUtilities.resolve_relative_path("ReadMe.md", codeunit_folder)
    GeneralUtilities.replace_regex_in_file(readme_file, "The latest version of PtN uses Go v\\d+\\.\\d+\\.\\d+\\.", f"The latest version of PtN uses Go v{latest_go_version}.")


if __name__ == "__main__":
    update_dependencies()
