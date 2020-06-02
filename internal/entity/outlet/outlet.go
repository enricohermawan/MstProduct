package outlet

import "gopkg.in/guregu/null.v3/zero"

// Outlet ...
type Outlet struct {
	OutRunningID        zero.String `db:"out_runningid" json:"out_runningid"`
	OutCode             zero.String `db:"out_code" json:"out_code"`
	OutName             zero.String `db:"out_name" json:"out_name"`
	OutAddress          zero.String `db:"out_address" json:"out_address"`
	OutCityCode         zero.Int    `db:"out_citycode" json:"out_citycode"`
	OutPostalCode       zero.String `db:"out_postalcode" json:"out_postalcode"`
	OutFax              zero.String `db:"out_fax" json:"out_fax"`
	OutEmail            zero.String `db:"out_email" json:"out_email"`
	OutPhoneunit        zero.Int    `db:"out_phoneunit" json:"out_phoneunit"`
	OutJnsarea          zero.Int    `db:"out_jnsarea" json:"out_jnsarea"`
	OutSquare           zero.Float  `db:"out_square" json:"out_square"`
	OutDateOpen         zero.String `db:"out_dateopen" json:"out_dateopen"`
	OutDateClose        zero.String `db:"out_dateclose" json:"out_dateclose"`
	OutTimeClose        zero.String `db:"out_timeclose" json:"out_timeclose"`
	OutNPWP             zero.String `db:"out_npwp" json:"out_npwp"`
	OutNPWPPPH          zero.String `db:"out_npwppph" json:"out_npwppph"`
	OutJenisOutlet      zero.Int    `db:"out_jenisoutlet" json:"out_jenisoutlet"`
	OutJenisKepemilikan zero.Int    `db:"out_jeniskepemilikan" json:"out_jeniskepemilikan"`
	OutPerjanjianYN     zero.String `db:"out_perjanjianyn" json:"out_perjanjianyn"`
	OutBuyingPower      zero.Float  `db:"out_buyingpower" json:"out_buyingpower"`
	OutClass            zero.String `db:"out_class" json:"out_class"`
	OutHOCode           zero.Int    `db:"out_hocode" json:"out_hocode"`
	OutAllowCreditYN    zero.String `db:"out_allowcredityn" json:"out_allowcredityn"`
	OutActiveYN         zero.String `db:"out_activeyn" json:"out_activeyn"`
	OutUserID           zero.String `db:"out_userid" json:"out_userid"`
	OutLastUpdate       zero.String `db:"out_lastupdate" json:"out_lastupdate"`
	OutDataAktifYN      zero.String `db:"out_dataaktifyn" json:"out_dataaktifyn"`
	OutJnsSTTK          zero.String `db:"out_jnssttk" json:"out_jnssttk"`
	OutJnsDeposit       zero.Int    `db:"out_jnsdeposit" json:"out_jnsdeposit"`
	OutJumpc            zero.Int    `db:"out_jumpc" json:"out_jumpc"`
	OutNamaPajak        zero.String `db:"out_nama_pajak" json:"out_nama_pajak"`
	OutAlamatPajak      zero.String `db:"out_alamat_pajak" json:"out_alamat_pajak"`
	OutScheduleRO       zero.String `db:"out_schedulero" json:"out_schedulero"`
	OutLuasGudang       zero.Float  `db:"out_luasgudang" json:"out_luasgudang"`
	OutLuasAreaJualan   zero.Float  `db:"out_luasareajualan" json:"out_luasareajualan"`
	OutCameraYN         zero.String `db:"out_camerayn" json:"out_camerayn"`
	OutAlamatPajakPPH   zero.String `db:"out_alamatpajakpph" json:"out_alamatpajakpph"`
	OutKPP              zero.String `db:"out_kpp" json:"out_kpp"`
	OutLokasiPajak      zero.String `db:"out_lokasipajak" json:"out_lokasipajak"`
	OutClassRenovasi    zero.String `db:"out_classrenovasi" json:"out_classrenovasi"`
	OutClassSpecial     zero.Int    `db:"out_classspecial" json:"out_classspecial"`
	OutJaminYN          zero.String `db:"out_jaminyn" json:"out_jaminyn"`
	OutJaminKet         zero.String `db:"out_jaminket" json:"out_jaminket"`
	OutRobotGDYN        zero.String `db:"out_robotgdyn" json:"out_robotgdyn"`
	OutTglPurgeRecord   zero.String `db:"out_tglpurgerecord" json:"out_tglpurgerecord"`
	OutTglArsipRecord   zero.String `db:"out_tglarsiprecord" json:"out_tglarsiprecord"`
}

// JSONTerima ...
type JSONTerima struct {
	Data Outlet `json:"data"`
}
