package config

import (
	"runtime"
	"sync"
)

var defaultConcurrency = runtime.GOMAXPROCS(0)

// Config 配置
type Config struct {
	InitPath    string // 初始化路径
	IgnoreHide  bool
	Debug       bool
	Exclude     []string // 排除文件夹,逗号分割
	SkipSuffix  []string // 跳过文件后缀
	Concurrency int
	Force       bool // 使用自定义配置覆盖默认配置
}

var c *Config
var once sync.Once

func init() {
	once.Do(func() {
		c = &Config{
			InitPath:   ".",
			IgnoreHide: true,
			Debug:      false,
			SkipSuffix: []string{".ico", ".tf", ".xsl",
				"gitkeep", ".log", ".po", ".pot", ".parquet",
				".dapper", ".json", ".3", ".2", ".1", ".000", ".tmp_drop", ".tmpfiles", ".sysusers", ".response", ".xcbuild", ".pbxproj",
				".3dm",
				".4ct",
				".4tc",
				".7z",
				".a",
				".a26",
				".a78",
				".aab",
				".aar",
				".acn",
				".acr",
				".agdai",
				".aif",
				".air",
				".alg",
				".ali",
				".aliases",
				".annot",
				".ap_",
				".api",
				".api-txt",
				".apk",
				".apng",
				".appx",
				".appxbundle",
				".appxupload",
				".aps",
				".asv",
				".atsuo",
				".autosave",
				".aux",
				".auxlock",
				".avi",
				".azurePubxml",
				".bak",
				".bbl",
				".bcf",
				".bck",
				".beam",
				".beams",
				".bgn",
				".bin",
				".binlog",
				".bit",
				".bld",
				".blg",
				".bmp",
				".booproj",
				".bowerrc",
				".box",
				".bpi",
				".bpl",
				".brf",
				".bridgesupport",
				".bs",
				".bsf",
				".byte",
				".bz2",
				".bzip",
				".bzip2",
				".c",
				".c_date",
				".ca",
				".cab",
				".cache",
				".cachefile",
				".cb",
				".cb2",
				".cfg",
				".cfgc",
				".changes",
				".chg",
				".class",
				".cma",
				".cmd",
				".cmd_log",
				".cmi",
				".cmo",
				".cmp",
				".cmx",
				".cmxa",
				".cmxs",
				".code-workspace",
				".compiled",
				".cover",
				".coverage",
				".coveragexml",
				".cpt",
				".crc",
				".crs",
				".crt",
				".csproj",
				".csr",
				".css.map",
				".csv",
				".ctxt",
				".cubin",
				".cut",
				".cycdx",
				".cyfit",
				".d",
				".d64fsl",
				".dat",
				".data",
				".db",
				".dbg",
				".dbmdl",
				".schemaview",
				".dbs",
				".dcp",
				".dcu",
				".deb",
				".debug",
				".def",
				".dep",
				".DEPLOYED",
				".der",
				".dex",
				".dfsl",
				".dib",
				".diff",
				".dll",
				".dmb",
				".dmg",
				".dmp",
				".dotCover",
				".dox",
				".dpth",
				".drc",
				".drd",
				".dres",
				".dri",
				".drl",
				".dsk",
				".dsn",
				".dSYM",
				".dta",
				".dump",
				".dvi",
				".dx32fsl",
				".dx64fsl",
				".dylib",
				".dyn_hi",
				".dyn_o",
				".e2e",
				".ear",
				".eep",
				".egg",
				".el",
				".elc",
				".elf",
				".eml",
				".end",
				".ent",
				".eps",
				".epub",
				".err",
				".errors",
				".esproj",
				".evcd",
				".eventlog",
				".ewt",
				".exe",
				".exe~",
				".exp",
				".ez",
				".FASL",
				".fatbin",
				".fdb_latexmk",
				".fff",
				".fls",
				".fmt",
				".fmx",
				".fot",
				".fsdb",
				".fx32fsl",
				".fx64fsl",
				".fyc",
				".gaux",
				".gcda",
				".gch",
				".gcno",
				".gcov",
				".gem",
				".ger",
				".gho",
				".gif",
				".gise",
				".glg",
				".glo",
				".glob",
				".gls",
				".glsdefs",
				".gph",
				".gpi",
				".gpState",
				".gpu",
				".gtex",
				".gz",
				".gzip",
				".h",
				".hdp",
				".hex",
				".hi",
				".hie",
				".hmap",
				".hp",
				".hpp",
				".hprof",
				".hst",
				".html",
				".hw",
				".hwdef",
				".i",
				".i*86",
				".ibc",
				".ico",
				".idb",
				".identcache",
				".idv",
				".idx",
				".ii",
				".ilg",
				".ilk",
				".image",
				".img",
				".img7",
				".iml",
				".ind",
				".info",
				".init",
				".int",
				".iobj",
				".ip_user_files",
				".ipa",
				".ipdb",
				".ipdefs",
				".iso",
				".itf",
				".iws",
				".j2k",
				".jar",
				".jfi",
				".jfif",
				".jfm",
				".jif",
				".cov",
				".mem",
				".jmconfig",
				".job",
				".jou",
				".jp2",
				".jpe",
				".jpeg",
				".jpf",
				".jpg",
				".jpm",
				".jpx",
				".js",
				".js_",
				".jsc",
				".jxr",
				".kdb",
				".kdev4",
				".kicad_pcb-bak",
				".kicad_prl",
				".kicad_sch-bak",
				".ko",
				".l15",
				".la",
				".lai",
				".launch",
				".lck",
				".lcov",
				".ldf",
				".lg",
				".lib",
				".lisp-temp",
				".listing",
				".lk",
				".ll",
				".llb",
				".lnk",
				".lo",
				".loa",
				".local",
				".lod",
				".loe",
				".lof",
				".log",
				".lol",
				".lot",
				".lox",
				".lps",
				".lpz",
				".lrs",
				".lrt",
				".lso",
				".lst",
				".ltjruby",
				".lvlibp",
				".lvlps",
				".lx32fsl",
				".lx64fsl",
				".ly2",
				".lyx#",
				".lyx~",
				".lzma",
				".lzo",
				".lzs",
				".m~",
				".maf",
				".manifest",
				".map",
				".max",
				".md5",
				".mdb",
				".mdf",
				".mem",
				".meta",
				".mex*",
				".mf",
				".mh",
				".mid",
				".midi",
				".mj2",
				".mlappinstall",
				".mlt",
				".mltbx",
				".mmx",
				".mng",
				".mno",
				".mo",
				".mobi",
				".moc",
				".mod",
				".mod*",
				".mode1v3",
				".mode2v3",
				".moved-aside",
				".mp",
				".mp3",
				".msd",
				".msi",
				".msix",
				".msk",
				".msm",
				".msp",
				".mti",
				".mw",
				".nar",
				".native",
				".nav",
				".ncb",
				".ncd",
				".ndf",
				".net",
				".ngc",
				".ngd",
				".ngr",
				".nlg",
				".nlo",
				".nls",
				".nupkg",
				".nvuser",
				".o",
				".obj",
				".obsolete",
				".ocx",
				".opendb",
				".opensdf",
				".opp",
				".opt",
				".opx",
				".broken",
				".or",
				".org",
				".ori",
				".orig",
				".os",
				".out",
				".p12",
				".p64fsl",
				".pad",
				".par",
				".patch",
				".pax",
				".pb",
				".pbxuser",
				".pcd",
				".pcf",
				".pch",
				".pdb",
				".pdb.meta",
				".pdf",
				".pdfpc",
				".pdfsync",
				".pem",
				".perspectivev3",
				".pfsl",
				".pfx",
				".pgc",
				".pgd",
				".pid",
				".lock",
				".pidb",
				".meta",
				".plc",
				".plg",
				".pls",
				".plt",
				".plx",
				".tdy",
				".png",
				".po",
				".pot",
				".ppk",
				".ppu",
				".pre",
				".prj",
				".prl",
				".pro",
				".prof",
				".projdata",
				".ps",
				".psess",
				".psm",
				".ptwx",
				".ptx",
				".publishproj",
				".publishsettings",
				".pubxml",
				".pxp",
				".pxt",
				".py,cover",
				".pyc",
				".pydevproject",
				".pyg",
				".pyo",
				".pytxcode",
				".ql",
				".qm",
				".qmlc",
				".qo",
				".rar",
				".raw",
				".rbb",
				".rbc",
				".rbd",
				".rbuistate",
				".rdb",
				".red",
				".rej",
				".rel",
				".res",
				".resources",
				".retry",
				".rkt.bak",
				".rkt~",
				".rnd",
				".rpm",
				".rpt",
				".rpx",
				".rsc",
				".rsj",
				".rsm",
				".rsp",
				".rst",
				".rsuser",
				".sav",
				".sbr",
				".scc",
				".sch-bak",
				".scm#*",
				".scm~",
				".sdf",
				".sdk",
				".seed",
				".ses",
				".sig",
				".skb",
				".sln",
				".sln.docstates",
				".slo",
				".slxc",
				".smcl",
				".smime",
				".sml",
				".smod",
				".snap",
				".snm",
				".snupkg",
				".so",
				".soc",
				".sol",
				".sout",
				".spec",
				".spl",
				".sql",
				".sqlite",
				".rock",
				".srec",
				".ss#*",
				".ss~",
				".ssleay",
				".sta",
				".stackdump",
				".stat",
				".stc",
				".stpr",
				".str",
				".sts",
				".stsem",
				".stx",
				".su",
				".sublime-workspace",
				".suo",
				".svclog",
				".svd",
				".svg",
				".svgz",
				".swf",
				".swp",
				".sx32fsl",
				".sx64fsl",
				".sym",
				".sympy",
				".synctex",
				".synctex.gz",
				".syr",
				".tar",
				".tar.gz",
				".tdo",
				".tds",
				".test",
				".tfm",
				".tfstate",
				".tfvars",
				".tgz",
				".thm",
				".tif",
				".tiff",
				".diff",
				".tlb",
				".tlh",
				".tli",
				".cache",
				".tmp",
				".tmp_proj",
				".tmproj",
				".tmproject",
				".toc",
				".tps",
				".translation",
				".trc",
				".tsbuildinfo",
				".tss",
				".ttt",
				".tvsconfig",
				".twr",
				".twx",
				".txz",
				".ucdb",
				".unitypackage",
				".unityproj",
				".unroutes",
				".upa",
				".upb",
				".user",
				".userosscache",
				".userprefs",
				".usertasks",
				".ut",
				".utf8.md",
				".uxp",
				".uxt",
				".after-timing",
				".before-timing",
				".d",
				".timing",
				".vbw",
				".db",
				".opendb",
				".vcd",
				".vdi",
				".vds",
				".veo",
				".ver",
				".vho",
				".vio",
				".vo",
				".vok",
				".vos",
				".vpd",
				".vpj",
				".vpw",
				".vpwhist",
				".vpwhistu",
				".vrb",
				".vsp",
				".vspscc",
				".vspx",
				".vssscc",
				".vstf",
				".vtg",
				".war",
				".wav",
				".wdf",
				".wdp",
				".webp",
				".wikidoc",
				".wlf",
				".wrt",
				".wx32fsl",
				".wx64fsl",
				".x86_64",
				".xar",
				".xbm",
				".xccheckout",
				".xcodeproj",
				".xcp",
				".xcscmblueprint",
				".xcworkspace",
				".xdv",
				".xdy",
				".xlk",
				".xln",
				".xml",
				".xmpi",
				".xojo_uistate",
				".xpi",
				".xref",
				".xst",
				".xwam",
				".xwm",
				".xyc",
				".xyd",
				".xz",
				".zip",
				".zo"},
			Exclude: []string{"node_modules", "vendor", "pod", "dist", "target",
				"bin", "asset", "img", ".vscode", ".idea", ".axoCover", "RECYCLE.BIN", "IntegrationServer", ".dSYM", ".egg-info", ".nuget", ".run", ".sln",
				"generated", "tmp", "Build", ".DesktopClient", ".bundle", ".eggs", ".cask", ".gradle", ".grunt", ".gwt", ".gwt-tmp", ".hypothesis", ".history", ".yarn",
				".import",
				".ionide",
				".jekyll-cache",
				".kdev4",
				".kitchen",
				".lein-plugins",
				".lgt_tmp",
				".localhistory",
				".metals",
				".mfractor",
				".mono",
				".mypy_cache",
				".navigation",
				".nb-gradle",
				".nox",
				".otto",
				".pybuilder",
				".pyre",
				".pytest_cache",
				".pytype",
				".recommenders",
				".rpt2_cache",
				".rts2_cache_cjs",
				".rts2_cache_es",
				".rts2_cache_umd",
				".sass-cache",
				".serverless",
				".settings",
				".stack-work",
				".svn",
				".temp",
				".tmp_versions",
				".tox",
				".vagrant",
				".vs",
				".anjuta",
				".build",
				".bundle",
				".fetch",
				".gitattributes",
				".grunt",
				".php_cs.cache",
				".cache",
				".web-server-pid",
				".yardoc",
				"_build",
				"_eumm",
				"_yardoc",
				"bazel-*",
				"bin",
				"bitrix",
				"blib",
				"cache",
				"cfg.py",
				"cfgcpp",
				"checkouts",
				"classes",
				"coverage",
				"CVS",
				"db",
				"depcomp",
				"deps",
				"dist",
				"log",
				"logs",
				"network-security.data",
				"node_modules",
				"package-cache",
				"play-cache",
				"play-stash",
				"py-compile",
				"stacktrace.log",
				"stamp-h1",
				"target",
				"temp",
				"test-build",
				"test-driver",
				"test-output",
				"test-report",
				"tmp",
				"typo3temp",
				"upload_backup",
				"vendor",
				"ylwrap",
				"__history",
				"__pycache__",
				"__pypackages__",
				"__recovery",
				"_checkouts",
				"_ngo",
				"_opam",
				"_ReSharper",
				"_site",
				"_thumbs",
				"_UpgradeReport_Files",
				"_xmsgs",
				"ASALocalRun",
				"AutoTest.Net",
				"BenchmarkDotNet.Artifacts",
				"Binaries",
				"bin-debug",
				"bin-release",
				"bld",
				"bower_components",
				"build_isolated",
				"build-iPhoneOS",
				"build-iPhoneSimulator",
				"BundleArtifacts",
				"compiled",
				"cython_debug",
				"data_*",
				"Debug",
				"demos",
				"Dependencies",
				"DerivedData",
				"DerivedDataCache",
				"devel",
				"devel_isolated",
				"develop-eggs",
				"fake-eggs",
				"FakesAssemblies",
				"gen",
				"Generated_Code",
				"gwt-unitCache",
				"Intermediate",
				"iOSInjectionProject",
				"jspm_packages",
				"lgt_tmp",
				"lib_managed",
				"logtalk_doclet_logs",
				"logtalk_tester_logs",
				"lvsRunDir",
				"MAlonzo",
				"MigrationBackup",
				"msg_gen",
				"out",
				"output",
				"override",
				"packer_cache",
				"paket-files",
				"pnacl",
				"Release",
				"sccprj",
				"sdist",
				"ServiceFabricBackup",
				"slprj",
				"src_managed",
				"srv_gen",
				"test-results",
				"test-servers",
				"textpattern",
				"thumbs",
				"urgReport",
				"web_modules",
				"workspaceuploads",
				"www-test",
				"xcuserdata",
				"xlnx_auto_0_xdb",
			},
			Concurrency: defaultConcurrency,
		}
	})
}

// SetConfig 配置Config
func (c *Config) SetConfig(ignoreHide bool, debug bool) *Config {
	c.IgnoreHide = ignoreHide
	c.Debug = debug
	return c
}

// GetInstance get an Instance
func GetInstance() *Config {
	return c
}
