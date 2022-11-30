package go_command

import (
	"fmt"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type queries interface {
	//SELECT * FROM @@table WHERE id=@id
	GetById(id int) (gen.T, error)
}

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
			Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		})

		gormdb, _ := gorm.Open(mysql.Open("jidan:jidan@(127.0.0.1:3306)/jsocial?charset=utf8mb4&parseTime=True&loc=Local"))

		g.UseDB(gormdb)

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
		tableModels := g.GenerateAllTable()
		for _, v := range tableModels {
			func(v interface{}) {
				g.ApplyInterface(func(queries) {}, v)
			}(v)
		}
		// execute the action of code generation
		g.Execute()

		fmt.Println("genGorm called")
	},
}

func init() {
	rootCmd.AddCommand(genGormCmd)
}
