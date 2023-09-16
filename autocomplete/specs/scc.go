// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["scc"] = model.Subcommand{
		Name:        []string{"scc"},
		Description: `Sloc, Cloc and Code. Count lines of code in a directory with complexity estimation`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "files or directories",
			IsOptional: true,
			IsVariadic: true,
		}},
		Options: []model.Option{{
			Name:        []string{"--avg-wage"},
			Description: `Average salary value used for COCOMO calculations`,
			Args: []model.Arg{{
				Name: "int",
			}},
		}, {
			Name:        []string{"--binary"},
			Description: `Disable binary file detection`,
		}, {
			Name:        []string{"--by-file"},
			Description: `Display output for every file`,
		}, {
			Name:        []string{"--ci"},
			Description: `Enable CI output settings where stdout is ASCII`,
		}, {
			Name:        []string{"--cocomo-project-type"},
			Description: `Change the COCOMO model type (allows custom models, eg. "name,1,1,1,1")`,
			Args: []model.Arg{{
				Name: "string",
				Suggestions: []model.Suggestion{{
					Name: []string{`organic`},
				}, {
					Name: []string{`semi-detached`},
				}, {
					Name: []string{`embedded`},
				}},
			}},
		}, {
			Name:        []string{"--count-as"},
			Description: `Count a file extension as a language (comma-separated key:value list, eg. jst:js,tpl:Markdown)`,
			Args: []model.Arg{{
				Name:      "string",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--debug"},
			Description: `Enable debug output`,
		}, {
			Name:        []string{"--exclude-dir"},
			Description: `Directories to exclude`,
			Args: []model.Arg{{
				Name:      "strings",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--file-gc-count"},
			Description: `Number of files to parse before turning the GC on`,
			Args: []model.Arg{{
				Name: "int",
			}},
		}, {
			Name:        []string{"-f", "--format"},
			Description: `Set output format`,
			Args: []model.Arg{{
				Name: "string",
				Suggestions: []model.Suggestion{{
					Name: []string{`tabular`},
				}, {
					Name: []string{`wide`},
				}, {
					Name: []string{`json`},
				}, {
					Name: []string{`csv`},
				}, {
					Name: []string{`csv-stream`},
				}, {
					Name: []string{`cloc-yaml`},
				}, {
					Name: []string{`html`},
				}, {
					Name: []string{`html-table`},
				}, {
					Name: []string{`sql`},
				}, {
					Name: []string{`sql-insert`},
				}},
			}},
		}, {
			Name:        []string{"--format-multi"},
			Description: `Multiple outputs with different formats (comma-separated key:value list, eg. tabular:stdout,csv:scc.csv)`,
			Args: []model.Arg{{
				Name:      "string",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--gen"},
			Description: `Identify generated files`,
		}, {
			Name:        []string{"--generated-markers"},
			Description: `Identify generated files by the presence of a string (comma-separated list)`,
			Args: []model.Arg{{
				Name: "strings",
			}},
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Help for scc`,
		}, {
			Name:        []string{"-i", "--include-ext"},
			Description: `Limit to these file extensions (comma-separated list)`,
			Args: []model.Arg{{
				Name:      "strings",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--include-symlinks"},
			Description: `Count symbolic links`,
		}, {
			Name:        []string{"-l", "--languages"},
			Description: `Print supported languages and extensions`,
		}, {
			Name:        []string{"--large-byte-count"},
			Description: `Number of bytes a file can contain before being omitted`,
			Args: []model.Arg{{
				Name: "int",
			}},
		}, {
			Name:        []string{"--large-line-count"},
			Description: `Number of lines a file can contain before being omitted`,
			Args: []model.Arg{{
				Name: "int",
			}},
		}, {
			Name:        []string{"--min"},
			Description: `Identify minified files`,
		}, {
			Name:        []string{"-z", "--min-gen"},
			Description: `Identify minified or generated files`,
		}, {
			Name:        []string{"--min-gen-line-length"},
			Description: `Number of bytes per average line for file to be considered minified or generated`,
			Args: []model.Arg{{
				Name: "int",
			}},
		}, {
			Name:        []string{"--no-cocomo"},
			Description: `Skip COCOMO calculation`,
		}, {
			Name:        []string{"-c", "--no-complexity"},
			Description: `Skip code complexity calculation`,
		}, {
			Name:        []string{"-d", "--no-duplicates"},
			Description: `Remove duplicate files from stats and output`,
		}, {
			Name:        []string{"--no-gen"},
			Description: `Ignore generated files in output (implies --gen)`,
		}, {
			Name:        []string{"--no-gitignore"},
			Description: `Disables .gitignore file logic`,
		}, {
			Name:        []string{"--no-ignore"},
			Description: `Disables .ignore file logic`,
		}, {
			Name:        []string{"--no-large"},
			Description: `Ignore files over certain byte and line size set by --max-line-count and --max-byte-count`,
		}, {
			Name:        []string{"--no-min"},
			Description: `Ignore minified files in output (implies --min)`,
		}, {
			Name:        []string{"--no-min-gen"},
			Description: `Ignore minified or generated files in output (implies --min-gen)`,
		}, {
			Name:        []string{"--no-size"},
			Description: `Remove size calculation output`,
		}, {
			Name:        []string{"-M", "--not-match"},
			Description: `Ignore files and directories matching regular expression`,
			Args: []model.Arg{{
				Name: "regex",
			}},
		}, {
			Name:        []string{"-o", "--output"},
			Description: `Output filename (defaults to stdout if not provided)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "string",
			}},
		}, {
			Name:        []string{"--remap-all"},
			Description: `Inspect every file and set its type by checking for a string (comma-separated key:value list, eg. "-*- C++ -*-":"C Header")`,
			Args: []model.Arg{{
				Name:      "string",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--remap-unknown"},
			Description: `Inspect files of unknown type and set its type by checking for a string (comma-separated key:value list, eg. "-*- C++ -*-":"C Header")`,
			Args: []model.Arg{{
				Name:      "string",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--size-unit"},
			Description: `Set the unit used for file size output`,
			Args: []model.Arg{{
				Name:        "string",
				Description: `See https://xkcd.com/394/`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`si`},
					Description: `1000^2 bytes`,
				}, {
					Name:        []string{`binary`},
					Description: `1024^2 bytes`,
				}, {
					Name:        []string{`mixed`},
					Description: `1,024,000 bytes (Binary kilobytes, SI megabytes)`,
				}, {
					Name:        []string{`xkcd-kb`},
					Description: `1000 bytes during leap years, 1024 otherwise`,
				}, {
					Name:        []string{`xkcd-kelly`},
					Description: `1012 bytes (a compromise between 1000 and 1024 bytes)`,
				}, {
					Name:        []string{`xkcd-imaginary`},
					Description: `1024*sqrt(-1) bytes (used in quantum computing)`,
				}, {
					Name:        []string{`xkcd-intel`},
					Description: `1023.937528 bytes (calculated on Pentium FPU)`,
				}, {
					Name:        []string{`xkcd-drive`},
					Description: `Currently 868 bytes (shrinks by 4 each year for marketing reasons)`,
				}, {
					Name:        []string{`xkcd-bakers`},
					Description: `1152 bytes (9 bits per byte, because you're such a good customer)`,
				}},
			}},
		}, {
			Name:        []string{"-s", "--sort"},
			Description: `Column to sort by`,
			Args: []model.Arg{{
				Name: "string",
				Suggestions: []model.Suggestion{{
					Name: []string{`files`},
				}, {
					Name: []string{`name`},
				}, {
					Name: []string{`lines`},
				}, {
					Name: []string{`blanks`},
				}, {
					Name: []string{`code`},
				}, {
					Name: []string{`comments`},
				}, {
					Name: []string{`complexity`},
				}},
			}},
		}, {
			Name:        []string{"--sql-project"},
			Description: `Use supplied name as the project identifier for the current run. Only valid with the '--format sql' or '--format sql-insert' option`,
			Args: []model.Arg{{
				Name: "string",
			}},
		}, {
			Name:        []string{"-t", "--trace"},
			Description: `Enable trace output (not recommended when processing multiple files)`,
		}, {
			Name:        []string{"-v", "--verbose"},
			Description: `Verbose output`,
		}, {
			Name:        []string{"--version"},
			Description: `Version for scc`,
		}, {
			Name:        []string{"-w", "--wide"},
			Description: `Wider output with additional statistics (implies --complexity)`,
		}},
	}
}