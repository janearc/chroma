# sources

every outside source the colourways, their samples and the spectra behind
them were drawn from. a picture is sampled and then not kept; what is kept
is numbers, and each kept file names its source again in its own header.

samples are in `samples/`. the spectra are kept beside the spectra tool, in
`cmd/spectra/data/`, until one of them becomes a colourway.

## pictures, sampled and not kept

- neptune, voyager 2. NASA/JPL, PIA01492 "Neptune Full Disk View", narrow
  angle camera, august 1989; green and orange filters only.
  https://science.nasa.gov/photojournal/neptune-full-disk-view
  kept as `samples/neptune-voyager.css`.
- pluto, 2015. NASA/JHUAPL/SwRI, PIA19708 "Pluto's Big Heart in Color",
  LORRI coloured from Ralph, 13 july 2015.
  https://science.nasa.gov/photojournal/plutos-big-heart-in-color
  kept as `samples/pluto-2015.css`.
- pluto, 2018. NASA/JHUAPL/SwRI/Alex Parker, "True Colors of Pluto", MVIC,
  14 july 2015, recalibrated to natural colour and released 23 july 2018.
  https://science.nasa.gov/resource/true-colors-of-pluto/ kept as
  `samples/pluto-2018.css`, and beside spectra as `data/pluto.css`.

## webb's neptune, spectra kept and the files not

from MAST, calibration level 3, fetched and deleted 2026-10-05.

- NIRSpec integral field unit, program 1249, 22 june 2023: two faces of
  the planet about eight hours apart, each through gratings G235H (F170LP)
  and G395H (F290LP), 1.66 to 5.27 microns, pipeline 2.0.1.

  `jw01249-o005_t001_nirspec_g235h-f170lp_s3d.fits`,
  `jw01249-o005_t001_nirspec_g395h-f290lp_s3d.fits`,
  `jw01249-o006_t001_nirspec_g235h-f170lp_s3d.fits`,
  `jw01249-o006_t001_nirspec_g395h-f290lp_s3d.fits`.

  kept as `neptune-nirspec-disc.csv`, `-clouds.csv` and `-clear.csv`:
  surface brightness in MJy/sr, the mean over a region, the two faces
  averaged.
- NIRCam, program 2739, 12 july 2022: four filters, F140M, F210M, F300M
  and F460M, pipeline 3.0.0.

  `jw02739-o004_t003_nircam_clear-f140m_i2d.fits` and the same for
  `f210m`, `f300m` and `f460m`.

  kept as `neptune-nircam-disc.csv`, `-clouds.csv` and `-clear.csv`: one
  value per filter, at its pivot wavelength.

the regions, in both: neptune's centre is where the header puts the
moving target (MT_RA, MT_DEC). the disc is the inner 85% of a 1.16
arcsecond radius, leaving out the limb.

clouds are the disc's brightest 15%, and clear its darker half, in
methane's band at 2.30 to 2.35 microns for NIRSpec and through F140M for
NIRCam, where only high cloud reflects.

the sky beside the planet is near zero in every file, so nothing is
subtracted.

triton was left out: its core saturated in NIRCam's short filters, and
its light came out about fifteen times too faint.

## spectra from the literature

- neptune's disc: Karkoschka, E. (1998), Icarus 133, 134-146,
  doi:10.1006/icar.1998.5913. NASA PDS Atmospheres Node, GBAT_0001,
  `DATA/1995LOW.TAB`. 300 to 1050 nm. kept as `neptune-disc.csv`.
- neptune's dark spot and bright clouds: Irwin, P. G. J. et al. (2023),
  Nature Astronomy 7, 1198-1207, doi:10.1038/s41550-023-02047-0. VLT MUSE;
  data at Zenodo, doi:10.5281/zenodo.7620656. kept as
  `neptune-dark-spot-*.csv` and `neptune-bright-clouds-sbs-difference.csv`.
- pluto's dark regions: Protopapa, S. et al. (2020), The Astronomical
  Journal 159, 74, doi:10.3847/1538-3881/ab5e82; VizieR J/AJ/159/74.
  kept as `pluto-cthulhu.csv`.

## models, not data

- gamma-ray bursts: the Band function with round values typical of the
  BATSE and Fermi catalogues, not any one burst. long: -1, -2.3, peak 200
  keV; short: -0.5, -2.3, peak 600 keV. behind `fermi-long.css` and
  `fermi-midnight.css`.
