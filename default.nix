{ buildGoPackage
}: buildGoPackage rec {
    pname = "twitch-cli";
    version = "alpha";

    src = ./.;

    goPackagePath = "go-theatron";
    goDeps = ./deps.nix;
}
