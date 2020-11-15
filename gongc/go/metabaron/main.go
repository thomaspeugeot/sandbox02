// the metabaron command explores a package and provides a visual
// representation
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/controllers"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
	"github.com/thomaspeugeot/sandbox02/gongc/go/walk"
)

const COMPUTED_FROM_PKG_PATH string = "computed from pkgPath (path to package for analysis)"

var (
	pkgPath = flag.String("pkgPath", "../../examples/bookstore/models", "path to package for analysis")

	backendTargetPath = flag.String("backendTargetPath", COMPUTED_FROM_PKG_PATH,
		"relative path to the directory where orm & controllers packages are generated"+
			" (by convention, one level above path to package for analysis)."+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	matTargetPath = flag.String("matTargetPath", COMPUTED_FROM_PKG_PATH, // "../../examples/bookstore/ng/src/app",
		"path to the ng directory where material components are generated"+
			"(by convention, relative to pkgPath, "+
			"into ../../ng/projects/<pkgName>/src/lib)"+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	ngWorkspacePath = flag.String("ngWorkspacePath", COMPUTED_FROM_PKG_PATH, // "../../examples/bookstore/ng",
		"path to the ng workspace directory (for performing npm isntall commands"+
			"(by convention, ../../ng relative to path to package for analysis)."+
			" If not set, it is "+COMPUTED_FROM_PKG_PATH)

	kamarRaimoPath = flag.String("kamarRaimoPath", "../../kamar-raimo",
		"realtive path to the directory containing templating files for api-api-generator")

	logBBFlag = flag.Bool("logDB", false, "log mode for db")

	skipSwagger = flag.Bool("skipSwagger", false, "skip swagger")

	loadAndGenerate = flag.Bool("loadAndGenerate", false, "parse package, generate code and exit")
	loaders         = flag.Bool("loaders", false, "generates loaders from a map of Struct into db")

	persist = flag.Bool("persist", false, "persist load & generate database into a metabaron.db file")

	addr = flag.String("addr", "localhost:8080",
		"network address addr where the angular generated service will lookup the server")

	apiFlag = flag.Bool("api", false, "it true, use api controllers instead of default controllers")
)

func main() {

	log.SetPrefix("metabaron: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// setup parse & generation parameters
	walk.RelativePkgPath = *pkgPath

	walk.ADDR = *addr

	// parse package and generate code if flag set
	if *loadAndGenerate {
		var db *gorm.DB
		if !*persist {
			db = models.SetupModels(*logBBFlag, ":memory:") // have in memory sqllite3 db
		} else {
			db = models.SetupModels(*logBBFlag, "metabaron.db")
		}
		walk.RemoveGoAllModelStruct(db)

		// load package into database
		walk.Walk(db)

		// compute path to generation directory
		{
			if *backendTargetPath == COMPUTED_FROM_PKG_PATH {
				*backendTargetPath = filepath.Join(*pkgPath, "..")
			}

			directory, err := filepath.Abs(*backendTargetPath)
			if err != nil {
				log.Fatal("Problem with backend target path " + err.Error())
			}

			// check existance of path
			fileInfo, err := os.Stat(directory)
			if os.IsNotExist(err) {
				log.Panicf("Folder %s does not exist.", directory)
			}
			if fileInfo.Mode().Perm()&(1<<(uint(7))) == 0 {
				log.Panicf("Folder %s is not writtable", directory)
			}
			walk.BackendTargetPath = directory
			log.Println("backend target path " + walk.BackendTargetPath)
		}
		{
			if *matTargetPath == COMPUTED_FROM_PKG_PATH {
				*matTargetPath = filepath.Join(*pkgPath, fmt.Sprintf("../../ng/projects/%s/src/lib", walk.PkgName))
			}

			directory, err := filepath.Abs(*matTargetPath)
			if err != nil {
				log.Panic("Problem with frontend target path " + err.Error())
			}
			// check existance of path
			fileInfo, err := os.Stat(directory)
			if os.IsNotExist(err) {
				log.Panicf("Folder %s does not exist.", directory)
			}
			if fileInfo.Mode().Perm()&(1<<(uint(7))) == 0 {
				log.Panicf("Folder %s is not writtable", directory)
			}
			walk.MatTargetPath = directory
			log.Println("module target abs path " + walk.MatTargetPath)
		}
		{
			if *ngWorkspacePath == COMPUTED_FROM_PKG_PATH {
				*ngWorkspacePath = filepath.Join(*pkgPath, "../../ng")
			}

			directory, err := filepath.Abs(*ngWorkspacePath)
			if err != nil {
				log.Panic("Problem with frontend target path " + err.Error())
			}
			walk.NgWorkspacePath = directory
			log.Println("module target abs path " + walk.NgWorkspacePath)
		}

		// generate directory for orm package
		walk.OrmPkgGenPath = filepath.Join(walk.BackendTargetPath, "orm")

		os.RemoveAll(walk.OrmPkgGenPath)
		errd := os.Mkdir(walk.OrmPkgGenPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + walk.OrmPkgGenPath)
		}
		if os.IsExist(errd) {
			log.Println("directory " + walk.OrmPkgGenPath + " allready exists")

			// supppress all files in it
		}

		if *apiFlag {
			// generate directory for api package
			walk.ApiPkgGenPath = filepath.Join(walk.BackendTargetPath, "api")

			os.RemoveAll(walk.ApiPkgGenPath)
			errd = os.Mkdir(walk.ApiPkgGenPath, os.ModePerm)
			if os.IsNotExist(errd) {
				log.Println("creating directory : " + walk.ApiPkgGenPath)
			}
			if os.IsExist(errd) {
				log.Println("directory " + walk.ApiPkgGenPath + " allready exists")

				// supppress all files in it
			}
		}

		// generate directory for controllers package
		walk.ControllersPkgGenPath = filepath.Join(walk.BackendTargetPath, "controllers")

		os.RemoveAll(walk.ControllersPkgGenPath)
		errd = os.Mkdir(walk.ControllersPkgGenPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + walk.ControllersPkgGenPath)
		}
		if os.IsExist(errd) {
			log.Println("directory " + walk.ControllersPkgGenPath + " allready exists")
		}

		// compute source path
		sourcePath, errd2 := filepath.Abs(walk.RelativePkgPath)
		if errd2 != nil {
			log.Panic("Problem with source path " + errd2.Error())
		}

		// generate files
		walk.GenORMSetup(db)

		// walk.GenJSONMarshallers(db)
		walk.GenORMModelDB(db, *loaders)
		walk.GenGoORMTranslation(db)
		walk.GenGoAllModelStruct(db)

		if *apiFlag {
			// new version
			walk.GenGoApiGate(db)
			walk.GenGoApiGateApi(db)
			walk.GenGoStructApi(db)

			// new version
			walk.GenControlersAPI(db)
		}

		walk.GenGoDocs(db)

		associationRoutesAndControllers := walk.GenControlers(db)
		walk.GenControllersRegistrations(db, associationRoutesAndControllers)
		walk.GenControllersErrorCodes(db)

		// generate code for mat library
		walk.GenNgSidebar(db)
		walk.GenNgTable(db)
		walk.GenNgRouting(db)
		walk.GenNgSplitter(db)
		walk.GenNgAdder(db)
		walk.GenNgDetail(db)
		walk.GenNgPresentation(db)
		walk.GenNgClass(db)
		walk.GenNgEnum(db)
		walk.GenNgService(db)
		walk.GenNgMatModuleApp(db)
		walk.GenNgIndex(db)

		apiYamlFilePath := fmt.Sprintf("%s/%sapi.yml", walk.ControllersPkgGenPath, walk.PkgName)
		if !*skipSwagger {

			// generate open api specification with swagger
			cmd := exec.Command("swagger",
				"generate", "spec",
				"-w", filepath.Dir(sourcePath), // swagger is interested in the "docs.go" in the package
				// by convention, this file is located on the parent dir
				"-o", apiYamlFilePath)
			log.Printf("Running command and waiting for it to finish...\n")

			// https://stackoverflow.com/questions/48253268/print-the-stdout-from-exec-command-in-real-time-in-go
			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			cmd.Stdout = mw
			cmd.Stderr = mw

			log.Println(cmd.String())
			log.Println(stdBuffer.String())

			// Execute the command
			if err := cmd.Run(); err != nil {
				log.Panic(err)
			}

		}
		os.Exit(0)
	}

	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := models.SetupModels(*logBBFlag, "test.db")

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	r.GET("/structs", controllers.GetStructs)
	r.GET("/structs/:id", controllers.GetStruct)
	r.POST("/structs", controllers.PostStruct)
	r.PATCH("/structs/:id", controllers.UpdateStruct)
	r.PUT("/structs/:id", controllers.UpdateStruct)
	r.DELETE("/structs/:id", controllers.DeleteStruct)

	r.GET("/fields", controllers.GetFields)
	r.GET("/fields/:id", controllers.GetField)
	r.POST("/fields", controllers.PostField)
	r.PATCH("/fields/:id", controllers.UpdateField)
	r.PUT("/fields/:id", controllers.UpdateField)
	r.DELETE("/fields/:id", controllers.DeleteField)

	r.GET("/diagrams", controllers.GetDiagrams)
	r.GET("/diagrams/:id", controllers.GetDiagram)
	r.POST("/diagrams", controllers.PostDiagram)
	r.PATCH("/diagrams/:id", controllers.UpdateDiagram)
	r.PUT("/diagrams/:id", controllers.UpdateDiagram)
	r.DELETE("/diagrams/:id", controllers.DeleteDiagram)

	r.POST("/actions", controllers.PostAction)

	r.Run()
}
