{
  description = "A basic gomod2nix flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix.url = "github:nix-community/gomod2nix";
    templ.url = "github:a-h/templ";
    templ.inputs.nixpkgs.follows = "nixpkgs";
    pre-commit-hooks.url = "github:cachix/pre-commit-hooks.nix";
    pre-commit-hooks.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = inputs:
    with inputs; let
      inherit (self) outputs;
    in
      flake-utils.lib.eachDefaultSystem (system: let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [gomod2nix.overlays.default];
        };
      in {
        packages.default = pkgs.buildGoApplication {
          pname = "myapp";
          version = "0.1";
          pwd = ./.;
          src = ./.;
          modules = ./gomod2nix.toml;
        };
        packages.goEnv = pkgs.mkGoEnv {pwd = ./.;};
        checks = {
          pre-commit-check = pre-commit-hooks.lib.${system}.run {
            src = ./.;
            hooks = {
              alejandra.enable = true;
              govet.enable = true;
              revive.enable = true;
              # staticcheck.enable = true;
            };
          };
        };
        devShells.default = pkgs.mkShell {
          inherit (self.checks.${system}.pre-commit-check) shellHook;
          packages = with pkgs; [
            git
            gnumake
            gomod2nix
            gopls
            gotools
            just
            go-tools
            inputs.templ.packages.${system}.templ
            outputs.packages.${system}.goEnv
          ];
        };
        formatter = pkgs.alejandra;
      });
}
