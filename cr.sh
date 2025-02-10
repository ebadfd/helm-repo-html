set -o errexit
set -o nounset
set -o pipefail

DEFAULT_CHART_RELEASER_VERSION=v0.0.1-hr-beta

show_help() {
  cat <<EOF
Usage: $(basename "$0") <options>

    -h, --help                    Display help
    -v, --version                 The helm-repo-html build version (default: $DEFAULT_CHART_RELEASER_VERSION)"
    -o, --owner                   The repo owner
    -r, --repo                    The repo name
        --pages-branch            The repo pages branch
    -t, --template                Template file
    -i, --input                   Input file
    -o, --output                  Ouptput file
        --use-arm                 Use ARM64 binary (default: false)
EOF
}

main() {
  local version="$DEFAULT_CHART_RELEASER_VERSION"
  local owner=
  local repo=
  local template=
  local input=
  local output=
  local pages_branch=
  local use_arm=false

  parse_command_line "$@"

  : "${CR_TOKEN:?Environment variable CR_TOKEN must be set}"

  local repo_root
  repo_root=$(git rev-parse --show-toplevel)
  pushd "$repo_root" >/dev/null

  popd >/dev/null
}

parse_command_line() {
  while :; do
    case "${1:-}" in
    -h | --help)
      show_help
      exit
      ;;
    -v | --version)
      if [[ -n "${2:-}" ]]; then
        version="$2"
        shift
      else
        echo "ERROR: '-v|--version' cannot be empty." >&2
        show_help
        exit 1
      fi
      ;;
    -t | --template)
      if [[ -n "${2:-}" ]]; then
        template="$2"
        shift
      else
        echo "ERROR: '-t|--template' cannot be empty." >&2
        show_help
        exit 1
      fi
      ;;
    -o | --owner)
      if [[ -n "${2:-}" ]]; then
        owner="$2"
        shift
      else
        echo "ERROR: '--owner' cannot be empty." >&2
        show_help
        exit 1
      fi
      ;;
    -r | --repo)
      if [[ -n "${2:-}" ]]; then
        repo="$2"
        shift
      else
        echo "ERROR: '--repo' cannot be empty." >&2
        show_help
        exit 1
      fi
      ;;
    --pages-branch)
      if [[ -n "${2:-}" ]]; then
        pages_branch="$2"
        shift
      fi
      ;;
    -i | --input)
      if [[ -n "${2:-}" ]]; then
        input="$2"
        shift
      fi
      ;;
    -o | --output)
      if [[ -n "${2:-}" ]]; then
        output="$2"
        shift
      fi
      ;;
    --use-arm)
      if [[ -n "${2:-}" ]]; then
          use_arm="$2"
          shift
      fi
      ;;
    *)
      break
      ;;
    esac

    shift
  done

  if [[ -z "$owner" ]]; then
    echo "ERROR: '-o|--owner' is required." >&2
    show_help
    exit 1
  fi

  if [[ -z "$repo" ]]; then
    echo "ERROR: '-r|--repo' is required." >&2
    show_help
    exit 1
  fi

  
  install_helm_repo_html()

  ls -lah
  helm-repo-html build -i $input -t $template -o $output

  git add .
  git commit -m 'core(hr): generated the home page'
  git push
}

install_helm_repo_html() {
  install_dir="$HOME/hr"

  mkdir -p "$install_dir"
  architecture=linux_amd64

  if [[ "$use_arm" = true ]]; then
    architecture=linux_arm64
  fi

  echo "Installing chart-releaser on $install_dir..."
  curl -sSLo cr.tar.gz "https://github.com/ebadfd/helm-repo-html/releases/download/$version/helm-repo-html_${version#v}_${architecture}.tar.gz"
  tar -xzf cr.tar.gz -C "$install_dir"
  rm -f cr.tar.gz

  echo 'Adding cr directory to PATH...'
  export PATH="$install_dir:$PATH"
}


main "$@"
