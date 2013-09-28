package main

import "flag"

func main() {
    pw := flag.Uint64("w", 40, "Width of the pictures of the colorscheme.")
    ph := flag.Uint64("h", 1024, "Height (number of levels) of the colorscheme.")
    flag.Parse()

    // No, I tested it, only using 3 or 4 of the keypoints isn't enough!
    colorschemes := map[string]GradientTable{
        // Sequential

        // Multihue
        "BuGn": GradientTable{
            {MustParseHex("#F7FCFD"), 0.000},
            {MustParseHex("#E5F5F9"), 0.125},
            {MustParseHex("#CCECE6"), 0.250},
            {MustParseHex("#99D8C9"), 0.375},
            {MustParseHex("#66C2A4"), 0.500},
            {MustParseHex("#41AE76"), 0.625},
            {MustParseHex("#238B45"), 0.750},
            {MustParseHex("#006D2C"), 0.875},
            {MustParseHex("#00441B"), 1.000},
        },
       "BuPu": GradientTable{
            {MustParseHex("#F7FCFD"), 0.000},
            {MustParseHex("#E0ECF4"), 0.125},
            {MustParseHex("#BFD3E6"), 0.250},
            {MustParseHex("#9EBCDA"), 0.375},
            {MustParseHex("#8C96C6"), 0.500},
            {MustParseHex("#8C6BB1"), 0.625},
            {MustParseHex("#88419D"), 0.750},
            {MustParseHex("#810F7C"), 0.875},
            {MustParseHex("#4D004B"), 1.000},
        },
        "GnBu": GradientTable{
            {MustParseHex("#F7FCF0"), 0.000},
            {MustParseHex("#E0F3DB"), 0.125},
            {MustParseHex("#CCEBC5"), 0.250},
            {MustParseHex("#A8DDB5"), 0.375},
            {MustParseHex("#7BCCC4"), 0.500},
            {MustParseHex("#4EB3D3"), 0.625},
            {MustParseHex("#2B8CBE"), 0.750},
            {MustParseHex("#0868AC"), 0.875},
            {MustParseHex("#084081"), 1.000},
        },
        "OrRd": GradientTable{
            {MustParseHex("#FFF7EC"), 0.000},
            {MustParseHex("#FEE8C8"), 0.125},
            {MustParseHex("#FDD49E"), 0.250},
            {MustParseHex("#FDBB84"), 0.375},
            {MustParseHex("#FC8D59"), 0.500},
            {MustParseHex("#EF6548"), 0.625},
            {MustParseHex("#D7301F"), 0.750},
            {MustParseHex("#B30000"), 0.875},
            {MustParseHex("#7F0000"), 1.000},
        },
        "PuBu": GradientTable{
            {MustParseHex("#FFF7FB"), 0.000},
            {MustParseHex("#ECE7F2"), 0.125},
            {MustParseHex("#D0D1E6"), 0.250},
            {MustParseHex("#A6BDDB"), 0.375},
            {MustParseHex("#74A9CF"), 0.500},
            {MustParseHex("#3690C0"), 0.625},
            {MustParseHex("#0570B0"), 0.750},
            {MustParseHex("#045A8D"), 0.875},
            {MustParseHex("#023858"), 1.000},
        },
        "PuBuGn": GradientTable{
            {MustParseHex("#FFF7FB"), 0.000},
            {MustParseHex("#ECE2F0"), 0.125},
            {MustParseHex("#D0D1E6"), 0.250},
            {MustParseHex("#A6BDDB"), 0.375},
            {MustParseHex("#67A9CF"), 0.500},
            {MustParseHex("#3690C0"), 0.625},
            {MustParseHex("#02818A"), 0.750},
            {MustParseHex("#016C59"), 0.875},
            {MustParseHex("#014636"), 1.000},
        },
        "PuRd": GradientTable{
            {MustParseHex("#F7F4F9"), 0.000},
            {MustParseHex("#E7E1EF"), 0.125},
            {MustParseHex("#D4B9DA"), 0.250},
            {MustParseHex("#C994C7"), 0.375},
            {MustParseHex("#DF65B0"), 0.500},
            {MustParseHex("#E7298A"), 0.625},
            {MustParseHex("#CE1256"), 0.750},
            {MustParseHex("#980043"), 0.875},
            {MustParseHex("#67001F"), 1.000},
        },
        "RdPu": GradientTable{
            {MustParseHex("#FFF7F3"), 0.000},
            {MustParseHex("#FDE0DD"), 0.125},
            {MustParseHex("#FCC5C0"), 0.250},
            {MustParseHex("#FA9FB5"), 0.375},
            {MustParseHex("#F768A1"), 0.500},
            {MustParseHex("#DD3497"), 0.625},
            {MustParseHex("#AE017E"), 0.750},
            {MustParseHex("#7A0177"), 0.875},
            {MustParseHex("#49006A"), 1.000},
        },
        "YlGn": GradientTable{
            {MustParseHex("#FFFFE5"), 0.000},
            {MustParseHex("#F7FCB9"), 0.125},
            {MustParseHex("#D9F0A3"), 0.250},
            {MustParseHex("#ADDD8E"), 0.375},
            {MustParseHex("#78C679"), 0.500},
            {MustParseHex("#41AB5D"), 0.625},
            {MustParseHex("#238443"), 0.750},
            {MustParseHex("#006837"), 0.875},
            {MustParseHex("#004529"), 1.000},
        },
        "YlGnBu": GradientTable{
            {MustParseHex("#FFFFD9"), 0.000},
            {MustParseHex("#EDF8B1"), 0.125},
            {MustParseHex("#C7E9B4"), 0.250},
            {MustParseHex("#7FCDBB"), 0.375},
            {MustParseHex("#41B6C4"), 0.500},
            {MustParseHex("#1D91C0"), 0.625},
            {MustParseHex("#225EA8"), 0.750},
            {MustParseHex("#253494"), 0.875},
            {MustParseHex("#081D58"), 1.000},
        },
        "YlOrBr": GradientTable{
            {MustParseHex("#FFFFE5"), 0.000},
            {MustParseHex("#FFF7BC"), 0.125},
            {MustParseHex("#FEE391"), 0.250},
            {MustParseHex("#FEC44F"), 0.375},
            {MustParseHex("#FE9929"), 0.500},
            {MustParseHex("#EC7014"), 0.625},
            {MustParseHex("#CC4C02"), 0.750},
            {MustParseHex("#993404"), 0.875},
            {MustParseHex("#662506"), 1.000},
        },
        "YlOrRd": GradientTable{
            {MustParseHex("#FFFFCC"), 0.000},
            {MustParseHex("#FFEDA0"), 0.125},
            {MustParseHex("#FED976"), 0.250},
            {MustParseHex("#FEB24C"), 0.375},
            {MustParseHex("#FD8D3C"), 0.500},
            {MustParseHex("#FC4E2A"), 0.625},
            {MustParseHex("#E31A1C"), 0.750},
            {MustParseHex("#BD0026"), 0.875},
            {MustParseHex("#800026"), 1.000},
        },
        // Single Hue
        "Blues": GradientTable{
            {MustParseHex("#F7FBFF"), 0.000},
            {MustParseHex("#DEEBF7"), 0.125},
            {MustParseHex("#C6DBEF"), 0.250},
            {MustParseHex("#9ECAE1"), 0.375},
            {MustParseHex("#6BAED6"), 0.500},
            {MustParseHex("#4292C6"), 0.625},
            {MustParseHex("#2171B5"), 0.750},
            {MustParseHex("#08519C"), 0.875},
            {MustParseHex("#08306B"), 1.000},
        },
        "Greens": GradientTable{
            {MustParseHex("#F7FCF5"), 0.000},
            {MustParseHex("#E5F5E0"), 0.125},
            {MustParseHex("#C7E9C0"), 0.250},
            {MustParseHex("#A1D99B"), 0.375},
            {MustParseHex("#74C476"), 0.500},
            {MustParseHex("#41AB5D"), 0.625},
            {MustParseHex("#238B45"), 0.750},
            {MustParseHex("#006D2C"), 0.875},
            {MustParseHex("#00441B"), 1.000},
        },
        "Greys": GradientTable{
            {MustParseHex("#FFFFFF"), 0.000},
            {MustParseHex("#F0F0F0"), 0.125},
            {MustParseHex("#D9D9D9"), 0.250},
            {MustParseHex("#BDBDBD"), 0.375},
            {MustParseHex("#969696"), 0.500},
            {MustParseHex("#737373"), 0.625},
            {MustParseHex("#525252"), 0.750},
            {MustParseHex("#252525"), 0.875},
            {MustParseHex("#000000"), 1.000},
        },
        "Oranges": GradientTable{
            {MustParseHex("#FFF5EB"), 0.000},
            {MustParseHex("#FEE6CE"), 0.125},
            {MustParseHex("#FDD0A2"), 0.250},
            {MustParseHex("#FDAE6B"), 0.375},
            {MustParseHex("#FD8D3C"), 0.500},
            {MustParseHex("#F16913"), 0.625},
            {MustParseHex("#D94801"), 0.750},
            {MustParseHex("#A63603"), 0.875},
            {MustParseHex("#7F2704"), 1.000},
        },
        "Purples": GradientTable{
            {MustParseHex("#FFF5EB"), 0.000},
            {MustParseHex("#FEE6CE"), 0.125},
            {MustParseHex("#FDD0A2"), 0.250},
            {MustParseHex("#FDAE6B"), 0.375},
            {MustParseHex("#FD8D3C"), 0.500},
            {MustParseHex("#F16913"), 0.625},
            {MustParseHex("#D94801"), 0.750},
            {MustParseHex("#A63603"), 0.875},
            {MustParseHex("#7F2704"), 1.000},
        },
        "Reds": GradientTable{
            {MustParseHex("#FFF5EB"), 0.000},
            {MustParseHex("#FEE6CE"), 0.125},
            {MustParseHex("#FDD0A2"), 0.250},
            {MustParseHex("#FDAE6B"), 0.375},
            {MustParseHex("#FD8D3C"), 0.500},
            {MustParseHex("#F16913"), 0.625},
            {MustParseHex("#D94801"), 0.750},
            {MustParseHex("#A63603"), 0.875},
            {MustParseHex("#7F2704"), 1.000},
        },

        // Diverging
        "BrBG": GradientTable{
            {MustParseHex("#543005"), 0.0},
            {MustParseHex("#8C510A"), 0.1},
            // The following contains a typo in green in colorbrewer2!
            {MustParseHex("#BFC22D"), 0.2},
            {MustParseHex("#DFC27D"), 0.3},
            {MustParseHex("#F6E8C3"), 0.4},
            {MustParseHex("#F5F5F5"), 0.5},
            {MustParseHex("#C7EAE5"), 0.6},
            {MustParseHex("#80CDC1"), 0.7},
            {MustParseHex("#35978F"), 0.8},
            {MustParseHex("#01665E"), 0.9},
            {MustParseHex("#003C30"), 1.0},
        },
        "PiYG": GradientTable{
            {MustParseHex("#8E0152"), 0.0},
            {MustParseHex("#C51B7D"), 0.1},
            {MustParseHex("#DE77AE"), 0.2},
            {MustParseHex("#F1B6DA"), 0.3},
            {MustParseHex("#FDE0EF"), 0.4},
            {MustParseHex("#F7F7F7"), 0.5},
            {MustParseHex("#E6F5D0"), 0.6},
            {MustParseHex("#B8E186"), 0.7},
            {MustParseHex("#7FBC41"), 0.8},
            {MustParseHex("#4D9221"), 0.9},
            {MustParseHex("#276419"), 1.0},
        },
        "PRGn": GradientTable{
            {MustParseHex("#40004B"), 0.0},
            {MustParseHex("#762A83"), 0.1},
            {MustParseHex("#9970AB"), 0.2},
            {MustParseHex("#C2A5CF"), 0.3},
            {MustParseHex("#E7D4E8"), 0.4},
            {MustParseHex("#F7F7F7"), 0.5},
            {MustParseHex("#D9F0D3"), 0.6},
            {MustParseHex("#A6DBA0"), 0.7},
            {MustParseHex("#5AAE61"), 0.8},
            {MustParseHex("#1B7837"), 0.9},
            {MustParseHex("#00441B"), 1.0},
        },
        "PuOr": GradientTable{
            {MustParseHex("#7F3B08"), 0.0},
            {MustParseHex("#B35806"), 0.1},
            {MustParseHex("#E08214"), 0.2},
            {MustParseHex("#FDB863"), 0.3},
            {MustParseHex("#FEE0B6"), 0.4},
            {MustParseHex("#F7F7F7"), 0.5},
            {MustParseHex("#D8DAEB"), 0.6},
            {MustParseHex("#B2ABD2"), 0.7},
            {MustParseHex("#8073AC"), 0.8},
            {MustParseHex("#542788"), 0.9},
            {MustParseHex("#2D004B"), 1.0},
        },
        "RdBu": GradientTable{
            {MustParseHex("#67001F"), 0.0},
            {MustParseHex("#B2182B"), 0.1},
            {MustParseHex("#D6604D"), 0.2},
            {MustParseHex("#F4A582"), 0.3},
            {MustParseHex("#FDDBC7"), 0.4},
            {MustParseHex("#F7F7F7"), 0.5},
            {MustParseHex("#D1E5F0"), 0.6},
            {MustParseHex("#92C5DE"), 0.7},
            {MustParseHex("#4393C3"), 0.8},
            {MustParseHex("#2166AC"), 0.9},
            {MustParseHex("#053061"), 1.0},
        },
        "RdGy": GradientTable{
            {MustParseHex("#67001F"), 0.0},
            {MustParseHex("#B2182B"), 0.1},
            {MustParseHex("#D6604D"), 0.2},
            {MustParseHex("#F4A582"), 0.3},
            {MustParseHex("#FDDBC7"), 0.4},
            {MustParseHex("#FFFFFF"), 0.5},
            {MustParseHex("#E0E0E0"), 0.6},
            {MustParseHex("#BABABA"), 0.7},
            {MustParseHex("#878787"), 0.8},
            {MustParseHex("#4D4D4D"), 0.9},
            {MustParseHex("#1A1A1A"), 1.0},
        },
        "RdYlBu": GradientTable{
            {MustParseHex("#A50026"), 0.0},
            {MustParseHex("#D73027"), 0.1},
            {MustParseHex("#F46D43"), 0.2},
            {MustParseHex("#FDAE61"), 0.3},
            {MustParseHex("#FEE090"), 0.4},
            {MustParseHex("#FFFFBF"), 0.5},
            {MustParseHex("#E0F3F8"), 0.6},
            {MustParseHex("#ABD9E9"), 0.7},
            {MustParseHex("#74ADD1"), 0.8},
            {MustParseHex("#4575B4"), 0.9},
            {MustParseHex("#313695"), 1.0},
        },
        "RdYlGn": GradientTable{
            {MustParseHex("#A50026"), 0.0},
            {MustParseHex("#D73027"), 0.1},
            {MustParseHex("#F46D43"), 0.2},
            {MustParseHex("#FDAE61"), 0.3},
            {MustParseHex("#FEE08B"), 0.4},
            {MustParseHex("#FFFFBF"), 0.5},
            {MustParseHex("#D9EF8B"), 0.6},
            {MustParseHex("#A6D96A"), 0.7},
            {MustParseHex("#66BD63"), 0.8},
            {MustParseHex("#1A9850"), 0.9},
            {MustParseHex("#006837"), 1.0},
        },
        "Spectral": GradientTable{
            {MustParseHex("#5e4fa2"), 0.0},
            {MustParseHex("#3288bd"), 0.1},
            {MustParseHex("#66c2a5"), 0.2},
            {MustParseHex("#abdda4"), 0.3},
            {MustParseHex("#e6f598"), 0.4},
            // {MustParseHex("#ffffbf"), 0.5},
            {MustParseHex("#fee090"), 0.6},
            {MustParseHex("#fdae61"), 0.7},
            {MustParseHex("#f46d43"), 0.8},
            {MustParseHex("#d53e4f"), 0.9},
            {MustParseHex("#9e0142"), 1.0},
        },
    }

    for name, keypoints := range colorschemes {
        CreateColorschemes(keypoints, name, int(*pw), int(*ph))
    }
}
