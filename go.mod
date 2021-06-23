module github.com/sergunya/go_mod_ch

go 1.16

require (
	github.com/akamensky/argparse v1.2.1
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/multipackage_1/package_2 v0.0.0-00010101000000-000000000000
)

replace github.com/multipackage_1/package_2 => ../multi_packages/package_2

exclude github.com/akamensky/argparse v1.2.2

retract v0.1.0
