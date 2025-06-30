{pkgs, ...}: {
  languages.go.package = pkgs.go_1_27;

  files.".golangci.yaml".copyMode = "seed";
}
