create database MyCampus;

create table DataMahasiswa (
    Id int primary key
    ,Nim float
    ,Nama_Mahasiswa varchar(300) not null
    ,Gender varchar(100) not null
    ,Jurusan varchar(100) not null
    ,Tahun_Ajaran float
    ,Jenis_Reg varchar(50) not null
    ,Tanggal_Masuk_Perkuliahan DATETIME
    ,Semester int
    ,Jumlah_Sks int
    ,Ipk DECIMAL
    ,Jumlah_MataKuliah_Lulus float
    ,Jumlah_MataKuliah_Td_Lulus float
    ,Jumlah_Sks_Lulus float
    ,Jumlah_Sks_Td_Lulus float
    ,Keterangan varchar(100) not null
)