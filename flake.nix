{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/e57b65abbbf7a2d5786acc86fdf56cde060ed026";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      rec {
        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              go_1_20
              nil
              nixpkgs-fmt
              dprint
              go-task
              goreleaser
            ];
          };

        packages.gh-action-multiline = pkgs.stdenv.mkDerivation
          {
            name = "gh-action-multiline";
            src = self;
            buildInputs = with pkgs; [
              go_1_20
              go-task
            ];
            buildPhase = ''
              # https://github.com/NixOS/nix/issues/670#issuecomment-1211700127
              export HOME=$(pwd)
              task build
            '';
            installPhase = ''
              mkdir -p $out/bin
              install -t $out/bin dist/gh-action-multiline
            '';
          };

        packages.default = packages.gh-action-multiline;

        # `nix run`
        apps.default = {
          type = "app";
          program = "${packages.gh-action-multiline}/bin/gh-action-multiline";
        };
      }
    );
}
