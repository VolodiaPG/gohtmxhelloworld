{
  description = "A basic gomod2nix flake";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";
  inputs.gomod2nix.url = "github:nix-community/gomod2nix";
  inputs.alejandra.url = "github:kamadorueda/alejandra/3.0.0";
  inputs.alejandra.inputs.nixpkgs.follows = "nixpkgs";
  inputs.templ.url = "github:a-h/templ";
  inputs.templ.inputs.nixpkgs.follows = "nixpkgs";

  outputs = inputs:
    with inputs; let
      inherit (self) outputs;
    in
      nixpkgs.lib.recursiveUpdate
      (flake-utils.lib.eachDefaultSystem (system: let
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
        devShells.default = pkgs.mkShell {
          packages = with pkgs;
            [
              git
              gomod2nix
              gopls
              gotools
              go-tools
              inputs.templ.packages.${system}.templ
              outputs.packages.${system}.goEnv
            ];
        };
      }))
      # Pass normal flake as `inputs` arg
      {
        formatter = alejandra.defaultPackage;
      };
}
