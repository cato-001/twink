{ pkgs }:

pkgs.buildGoApplication {
  pname = "twink";
  version = "0.0.1";
  src = ./.;
  modules = ./gomod2nix.toml;
}
