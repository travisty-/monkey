{ pkgs, ... }:

{
  # https://devenv.sh/basics
  env.SHELL = "${pkgs.zsh}/bin/zsh";

  # https://devenv.sh/packages
  packages = [
    pkgs.git
    pkgs.zsh
  ];

  # https://devenv.sh/languages
  languages.go.enable = true;

  # https://devenv.sh/scripts
  enterShell = ''
    echo "• Entering shell ..."
  '';

  # https://devenv.sh/git-hooks
  git-hooks.hooks = {
    gofmt.enable = true;
    golangci-lint.enable = true;
    golines.enable = true;
    gotest.enable = true;
    govet.enable = true;
    staticcheck.enable = true;
  };
}
