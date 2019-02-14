// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "78aa01be528be8fc333fc794bc4a3c1b"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"4526cfd1d77bab7843d30700c2868275": "1f8b08000000000000ffac94516bdb400cc7dffb29340f460b251d648c61285bb61656284d4952c69ee28bad24c7ce2757a74b1742befb909d3849b70c56f69060cb27e92ffda42b70818eaa12bda4270085350e7349212997e1d1256a32622626600ac96ad5a14a42e79e716a7faed7e33d6f3d3aa7a0ae8e72e3f4596d15b1dabaddb7eff53506e41412266abe9a109e888bd67422188e2979ad3ff84e1172e335121808d6cf1cc2c3e01682b0f533981283cc1182291172f2533b8b6cc4924fb73122bb6ddcf4e24213a7faf7e9b4559eaadeb38bdf0a56751f2bc30147b6c44be1886fcae8c40ec5086a1f42636434851ea12897ddb0537f478220732360a7b0a45897914576d93990cc910f2507308c10ac432f6e0976e689b1e8d4a16e3cc8dc06c84dc0f3ba64ed421983680c31d683710e181fa3652cd4e831d7a8501936250a72e86c94fd8db216fdbff0564c45cc0f68b4448d0b54f703fdc232796da771b0306ccdc4610021a00532db0261615cc40076db84ba6b9d03c2ab55b25a25eb35fa459f21b9ea8d7a9f7bc3ebf1c3e0367919fe9dfa970cc17abd6df6bf6dd92eeb8ec2b1dabef687a3649f519db4a174cce7be3f509f9a617dbca178b47dc3eb41b201da446f291fcdd01b0ebff50757076e0dfc9e2fea3d6807a025fcc77985c912028ae8a65b81e80b64a0aade9596fee6bd9e3095b721958252698ccf78ed7fdaa39642373446f43915d6cf5248a24c3f949377e3197a64e3c6b9ddadf797dec3e8a67ff7aa7182accd9d81f145f63c6dd6aceb04b5261def4c6564faf444fc032aa60ad92d3b9b7837cd9da1778227811019cfa1a0fa259f1b3f433825d660677a2184ed9a744e7e050000ffff79f0e5cfe8050000",
		"53687fd7c4aa5237e9bf81249c4bd931": "1f8b08000000000000ffecd1416b32311006e0fbfe8a21df7d51c14fd89b6c05c1426463293d494ca6b8189d25136dcb92ff5e62517b28dbe24df03ae4c99bc96bf1808e9a2dee429101d85a3b34a1004366e3499b751aeaa0579ab180b6cda9099ccf3dbed6ef312ebff10c604d1cd221d1b62246c0dd417a10a52c67951c97d3e554aa8500d11f8cf25edecbfb0220c60ca021dfe9e6b24a6ef07f301c9dcc9ed177992735a904084f14ce319af98dbced8c1a2bf52cab0701e2f23a72050c33006a424d3b4e1f05c0ecb664b1005bb35e393c0eff31bb1469b07b23a51e9795948b7272dcec14f575c15ff08f70831fbfb9d9e4e5ccb2807c55ebc9ddebbeb9ba1b4f766fd24ed7947ed1f7ea6fadfacf000000ffff1757d2dbe6050000",
		"793e1955dae01f2888949d22a97bc7af": "1f8b08000000000000ffac8ec14a03311086ef798a210f906dab450888aed89ba0d8f5bca4cda80be94ec8cc5621e4dd2542ab45bd79fbf9997fbecfe31e03c51d8e6215801f5cc0ad5888c4f292906be7c46d1ca3859c0d4561f390f079782fa5ffb6560013633a5946c7fc46c99f94afc46261beb830333333f37a46142c2c9512e44f8b29850ad339eb5270dcdf27d0dd6addf5b76dd7deb4eb55fff478a7411fbedaa639c643b83e12ecf2fc6cd1fc70afb02be6b0238f977e60b709a84b512a26f2d356061aff72f94f8d2fda6f321f010000ffffb44a8cd79f010000",
		"d6126d474cf9b7af331e89fba0c6f146": "1f8b08000000000000ff4a492d4bcdc92fc84dcd2bb1e2525048c94ccc494d2eb152502a2eccc92c4935560209269624262516a75a295457ebe5179414eb05e5e797d4d6c6c3b80145a9699915b5b5f148a6e9410ce0e22a492da686d12063e0661614e5a794269764e6e751c164846130f301010000ffffd5f1ef6a15010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("pop:genny:config", "../config/templates")
		b.SetResolver("cockroach.yml.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "53687fd7c4aa5237e9bf81249c4bd931"})
		b.SetResolver("mysql.yml.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "4526cfd1d77bab7843d30700c2868275"})
		b.SetResolver("postgres.yml.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "793e1955dae01f2888949d22a97bc7af"})
		b.SetResolver("sqlite3.yml.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "d6126d474cf9b7af331e89fba0c6f146"})
	}()

	return nil
}()
