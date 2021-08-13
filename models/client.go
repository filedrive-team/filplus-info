package models

import (
	"github.com/filedrive-team/filplus-info/settings"
	"github.com/filedrive-team/filplus-info/types"
	"github.com/filedrive-team/filplus-info/utils"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type Client struct {
	Model
	Name    string `json:"name" gorm:"size:64;uniqueIndex:uix_client_name_addr"`     //client name
	Website string `json:"website" gorm:"size:255"`                                  //website、social media
	Region  string `json:"region" gorm:"size:64"`                                    //region
	Address string `json:"address" gorm:"size:255;uniqueIndex:uix_client_name_addr"` //client address
}

func (Client) TableName() string {
	return "client"
}

func init() {
	autoMigrateModels = append(autoMigrateModels, &Client{})
}

func GetClientList(name, address string, offset, size int) (total int, list []*Client, err error) {
	list = make([]*Client, 0)
	stmt := db.Model(Client{})
	if name != "" {
		stmt = stmt.Where("name=?", name)
	}
	if address != "" {
		stmt = stmt.Where("address=?", address)
	}
	var tmpTotal int64
	err = stmt.Count(&tmpTotal).Error
	if err != nil {
		err = errors.Wrap(err, "count client table failed.")
		return
	}
	total = int(tmpTotal)
	err = stmt.Order("id desc").Offset(offset).Limit(size).Find(&list).Error
	if err != nil {
		err = errors.Wrap(err, "query client table failed.")
		return
	}
	return
}

func ClientList(clientName string, client string, notary []string, params *types.PaginationParams) (*types.CommonList, error) {
	type ClientItem struct {
		Address    string          `json:"address"`
		Client     string          `json:"client"`
		ClientName string          `json:"client_name"`
		Allowance  decimal.Decimal `json:"allowance"`
		BlockTime  int64           `json:"block_time"`
		Epoch      uint64          `json:"epoch"`
		NotaryName string          `json:"notary_name"`
	}

	offset, size := utils.PaginationHelper(params.Page, params.PageSize, settings.DefaultPageSize)

	stmtCount := `
SELECT 
count(*)
from client_allowance ca 
left join client c
on ca.client = c.address 
where c.deleted_at is null
`
	stmt := `
SELECT 
ca.notary as address,
ca.allowance,
ca.client,
COALESCE(c.name,'') as client_name,
ca.block_time,
ca.epoch, 
n.notary_name 
from client_allowance ca 
left join client c
on ca.client = c.address 
left join notary n
on ca.notary = n.address 
where c.deleted_at is null
`
	var stmtParams []interface{}
	if len(notary) > 0 {
		stmtCount += " and ca.notary in ?"
		stmt += " and ca.notary in ?"
		stmtParams = append(stmtParams, notary)
	}
	if clientName != "" {
		stmtCount += " and c.name=?"
		stmt += " and c.name=?"
		stmtParams = append(stmtParams, clientName)
	}
	if client != "" {
		stmtCount += " and ca.client=?"
		stmt += " and ca.client=?"
		stmtParams = append(stmtParams, client)
	}
	total := 0
	err := db.Debug().Raw(stmtCount, stmtParams...).Scan(&total).Error
	if err != nil {
		return nil, err
	}
	stmt += " order by ca.block_time desc limit ? offset ?"
	stmtParams = append(stmtParams, size, offset)
	var data []*ClientItem
	err = db.Debug().Raw(stmt, stmtParams...).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	res := &types.CommonList{
		Total: total,
		List:  data,
	}
	return res, nil
}

func InsertClient(c *Client) error {
	return db.Create(c).Error
}

func UpsertClient(c *Client) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "address"}},
		DoUpdates: clause.AssignmentColumns([]string{"website", "region"}),
	}).Create(c).Error
}

func TruncateClient() error {
	return db.Exec("truncate client;").Error
}

func InitClientList() error {
	var total int64
	err := db.Model(Client{}).Count(&total).Error
	if err != nil {
		err = errors.Wrap(err, "count client table failed.")
		return err
	}
	if total > 0 {
		logger.Info("client table is not empty, skip it")
		return nil
	}
	list := []*Client{
		{
			Name:    "Alyale Philippines",
			Address: "f3qqhbwhgorbtymzhq5wxnointwl4tl5riqvxv35utyuhlisbz4247io32oshff2emysu5fzc2t4msh64ahlpq",
		},
		{
			Name:    "Askender",
			Address: "f3uvwjknloel3zcmjrgrbmgqwp3c32l7b7tyoqj7p7njb3c2xmex6dh2jzz45a6smzqvhw53nbng2zn2p7ciza",
		},
		{
			Name:    "ElioVP",
			Address: "f3rnzgvmh6itjywfy3n6e23s3impootcafbook4x3et5qqmxrces6hd5oypufdwkwi6eobnhfqny3oohr2s5ea",
		},
		{
			Name:    "Filbox",
			Address: "f3umho76jhnreafz5nraeyob25kidea26xz5xzewpe5glc45ullz55bzusfnfrk2nlb3ibd3aiphhmblqpuvda",
		},
		{
			Name:    "FileDrive",
			Address: "f3wgfwtrs5p6jrkwfl2mksqa2ivgbgdjjrhjbefy3n7qzvotc3y6sazmp5gfyj7um6jlgdvlbiepzawnc6wxtq",
		}, {
			Name:    "FileDrive",
			Address: "f3s3dwh67ps4kxui4mziwvus5ue7p6p324aiebrrpe24st5t334ib5gwlxql6auo7ibrccmppjxxavah46wdqa",
		},
		{
			Name:    "Guazi Dynamic",
			Address: "f3qaipvxrz2gxexc7mcvjsxifmscfiw7c7zfhrmq76j5ee4hcbvg3gbtpea7wgz72kkjcjmwzhm5uo2onxyocq",
		},
		{
			Name:    "Guazi Dynamic",
			Address: "f3wvmap64wokgkvavberi2l3o7li5egvym526kdgk3bmnmto5qbcq7jin3wpbkiifoa26caxreois237d66nrq",
		},
		{
			Name:    "IPFS Force",
			Address: "f3rtqrpxadmhkfgy3cmtib4q6pwiqxpn6fxxq7kvopwqtwghwhwhyy7lm4euqvkg5h26ll3fpc54zsv7jtwdna",
		},
		{
			Name:    "IPFS Union",
			Address: "f3rf334djo4e76f44hz6tbkh5maerg42nr5b5squsb7kty336zv6pdnnwhyc2qaijpq2dwkpyy3fwfeofjtpja",
		},
		{
			Name:    "JV",
			Address: "f1s566sc7nmxo7b7qml7m6s72uwmljg6k5orliolq",
		},
		{
			Name:    "Kernelogic",
			Address: "f3rfb7aarwmrulwc4rqy4526zvopnsnbsmcfokhbi743kq2shaybo46cbjqezvmhhjkejpk6xqbafixvplq5xa",
		},
		{
			Name:    "Kernelogic",
			Address: "f3viailjjkez5o5d3p2flxkdkag2srqy3wouaxcemabrqjp2nr6l56yxbx5pu3qgszgo7dgrzpon4xpuweny5a",
		},
		{
			Name:    "LandSatFile",
			Address: "f3rikexkas2vxfhc6nbhla7hpunrp6as5msjx3wwndve6rfgzmyrliu2p3ypzdf74yzfwjhyqpnt6gj4xwzkoa",
		},
		{
			Name:    "Linix Webhosting",
			Address: "f1jqjoleqxffusnsyxxlmyedhqsiigoqg2eg4sgzq",
		},
		{
			Name:    "NBFS",
			Address: "f3rkcl5lnoumsunl3grsy2b57z7zt6uoyvgnviz2dw3ir7sslnya2634as2ksfxzziuxzgubavqsb3s2fmofcq",
		}, {
			Name:    "NBFS",
			Address: "f3ufzpudvsjqyiholpxiqoomsd2svy26jvy4z4pzodikgovkhkp6ioxf5p4jbpnf7tgyg67dny4j75e7og7zeq",
		},
		{
			Name:    "Neo Beat",
			Address: "f3uqb2gqegru5ycchufy6gyjbpiu3dk3tx4nk45aua6wd6krbvkxyj6qapu6kkicm7nyzpdpg22clgwwy2wwba",
		},
		{
			Name:    "Phi Mining",
			Address: "f3v2bujmjhxugk3wtwwqha3za7vxmetyvd5rfj2dxff7zekibcclmbhd2ytrmrfls7bb2bz53ehcpl5npss5hq",
		},
		{
			Name:    "Piñata",
			Address: "f17kamajrdzjcjzj6b3y2ovc64kc7qxvkbx2dzvza",
		},
		{
			Name:    "Shanghai Futures Exchange",
			Address: "f1jp2xt7qkgdfmsj6pylcalmr7t4atfby2aoyeh4i",
		},
		{
			Name:    "Shenzhen Xiaoyu Blockchain Technology",
			Address: "f1ry5tgrth2qt4r5zu4e6ssfkf4qrce7hcdp22kbi",
		},
		{
			Name:    "Simba",
			Address: "f3srltdd7fvnukjpywap5wksbpmdenhvwr6imza3nxlxycmmi2lit2kbisij5wly5fmjwcatmetyktpd5tfydq",
		},
		{
			Name:    "Speedium",
			Address: "f1luelfcqktlgpw55cqsqieyilusqrayzxbctsqfa",
		},
		{
			Name:    "Speedium",
			Address: "f3qfjtedox66c3x6dgqt45e364ht3ydcfymcx6xucoy34s22y6jttatbcatsuppesupanb5476npnt5tdasrqa",
		},
		{
			Name:    "Speedium",
			Address: "f3s5qozdyocudxtt2ekvcbu34dj2ubt3b4fqdzjqla7gj7442xckmlntcsmzibru6zkpiqlwflgn554cdhw6va",
		},
		{
			Name:    "Speedium",
			Address: "f3sbunigqcrdk2xkxcgvrm72wt73fb4km2kykuy3xvb5gexf66kryjoweno2zlz6kyjuzcket72xd2s7b3zhoq",
		},
		{
			Name:    "Speedium",
			Address: "f3shjxrnqvzn6x2dxojjpfplq5hhxww4qdtnfm42wosayutnjhtahpeu3qjcidz6y7kmqjrw6uuahhiua4vqyq",
		},
		{
			Name:    "Speedium",
			Address: "f3tc2ebpyy7uojoewaxjec3hmsh4dmrdxdgmeyemnzr4hvwr7a7x5nvr5ykbjmsfaclychm5dha53inaiyq3da",
		},
		{
			Name:    "Speedium",
			Address: "f3tf2q7aimiuezmlfe2obixtyjty3atxbwdz3vaqiekx3j3styocv6zgebkwdycpk2j5bvzstmosuqziezej2a",
		},
		{
			Name:    "Starry sky in Yunnan",
			Address: "f3qqttrr5b6df5kjrtt6fwy3bafthqboe3aeghhqas2qiwavlxjlbdmo5xsai4s5hnam53pdiq55wsvkf52s6q",
		},
		{
			Name:    "Wang Liang",
			Address: "f3sar6tat2mhqadlqjjxewmpqya3whpkzuqaevpo5vwfvdhnm3uztd7p67d6ozpsfkiwvrvqrhp3zu6psi23ma",
		},
		{
			Name:    "Wang Liang",
			Address: "f3vsg65k7fs365uueavmfnwcd2wzxq62vdhh7ydgws6oi2xliwjenyp75cg4ynxjez3i2migafpgv26iyee74q",
		},
		{
			Name:    "liyiping",
			Address: "f1jy3blozja7edk7firimlylpxpqkqzrx34grkcgy",
		},
		{
			Name:    "s0nik42",
			Address: "f3waqhsynu4a3w2a4mcwfaprd4fu6rgd3uuvt7mwj5v24br7u6vcb37xbvgbdgekaav3vcc3m6spz6etuh4aja",
		},
		{
			Name:    "上海小工蚁电子商务股份有限公司",
			Address: "f1xl5b4vmjnikoku4j6brt33yoyqqepmqt7aowzny",
		},
		{
			Name:    "北京中电博顺智能设备技术有限公司",
			Address: "f1uhubdateuntjmpje6yg4dwy6p5hjwxewuddueiy",
		},
		{
			Name:    "安徽六六六科技有限公司",
			Address: "f14ew4fybr3grirbaii2yj4mmfo77kpayu74xegbi",
		},
		{
			Name:    "杭州讯酷科技有限公司",
			Address: "f1wninfaoogdihtbqntlhjanb3z5bswqvjcestuuq",
		},
	}
	for _, c := range list {
		err := InsertClient(c)
		if err != nil {
			return err
		}
	}
	return nil
}
