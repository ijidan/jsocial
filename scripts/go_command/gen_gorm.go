package go_command

import (
	"fmt"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

// genGormCmd represents the genGorm command
var genGormCmd = &cobra.Command{
	Use:   "gen:gorm",
	Short: "auto gen gorm",
	Long:  `auto gen gorm`,
	Run: func(cmd *cobra.Command, args []string) {

		rootPath := Path.ProjectDir
		configFile := Path.ConfigsDir + "api.yaml"
		param := config.NewCmdParam(rootPath, configFile)
		global.GH, _ = injector.InitializeEventHttpGlobal(*param)

		// specify the output directory (default: "./query")
		// ### if you want to query without context constrain, set mode gen.WithoutContext ###
		g := gen.NewGenerator(gen.Config{
			OutPath: Path.InternalDir + "query",
			/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
			//if you want the nullable field generation property to be pointer type, set FieldNullable true
			/* FieldNullable: true,*/
			//if you want to generate index tags from database, set FieldWithIndexTag true
			/* FieldWithIndexTag: true,*/
			//if you want to generate type tags from database, set FieldWithTypeTag true
			/* FieldWithTypeTag: true,*/
			//if you need unit tests for query code, set WithUnitTest true
			/* WithUnitTest: true, */
		})

		// reuse the database connection in Project or create a connection here
		// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
		// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
		g.UseDB(global.GH.Db)

		// apply basic crud api on structs or table models which is specified by table name with function
		// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
		//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))

		// apply diy interfaces on structs or table models
		//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

		device := g.GenerateModel("device")
		deviceAck := g.GenerateModel("device_ack", gen.FieldRelate(field.BelongsTo, "Device", device, &field.RelateConfig{
			// RelateSlice: true,
			GORMTag: "foreignKey:device_id",
		}))
		messageIndex := g.GenerateModel("message_index")
		groupUser := g.GenerateModel("group_user")

		user := g.GenerateModel("user",
			gen.FieldRelate(field.HasMany, "Device", device, &field.RelateConfig{
				// RelateSlice: true,
				GORMTag: "foreignKey:device_id",
			}),
			gen.FieldRelate(field.HasMany, "MessageIndex", device, &field.RelateConfig{
				// RelateSlice: true,
				GORMTag: "foreignKey:receiver_id",
			}),
			gen.FieldRelate(field.Many2Many, "GroupUser", groupUser, &field.RelateConfig{
				// RelateSlice: true,
				GORMTag: "foreignKey:user_id",
			}),
		)
		group := g.GenerateModel("group", gen.FieldRelate(field.Many2Many, "GroupUser", groupUser,
			&field.RelateConfig{
				// RelateSlice: true,
				GORMTag: "foreignKey:group_id",
			}))

		g.ApplyBasic(device, deviceAck, group, groupUser, messageIndex, user)
		g.ApplyBasic(g.GenerateAllTable()...)
		// execute the action of code generation
		g.Execute()

		fmt.Println("genGorm called")
	},
}

func init() {
	rootCmd.AddCommand(genGormCmd)
}
