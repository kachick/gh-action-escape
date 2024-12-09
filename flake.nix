{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.11";
    selfup = {
      url = "github:kachick/selfup/v1.1.8";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    {
      self,
      nixpkgs,
      selfup,
    }:
    let
      inherit (nixpkgs) lib;
      forAllSystems = lib.genAttrs lib.systems.flakeExposed;
    in
    {
      formatter = forAllSystems (system: nixpkgs.legacyPackages.${system}.nixfmt-rfc-style);
      devShells = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShellNoCC {
            buildInputs =
              (with pkgs; [
                # https://github.com/NixOS/nix/issues/730#issuecomment-162323824
                # https://github.com/kachick/dotfiles/pull/228
                bashInteractive
                findutils # xargs
                nixfmt-rfc-style
                nil

                go_1_23
                go-task
                goreleaser

                dprint
                typos
              ])
              ++ [ selfup.packages.${system}.default ];
          };
        }
      );

      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.stdenv.mkDerivation {
            name = "gh-action-escape";
            src = self;
            buildInputs = with pkgs; [
              go_1_23
              go-task
            ];
            buildPhase = ''
              # https://github.com/NixOS/nix/issues/670#issuecomment-1211700127
              export HOME=$(pwd)
              task build
            '';
            installPhase = ''
              mkdir -p $out/bin
              install -t $out/bin dist/bin/gh-action-escape
            '';
          };
        }
      );

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = nixpkgs.lib.getExe self.packages.${system}.default;
        };
      });
    };
}
