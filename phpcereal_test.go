package phpcereal_test

import (
	"testing"

	"github.com/mikeschinkel/go-phpcereal"
	"github.com/stretchr/testify/assert"
)

//const (
//	sAll      = `a:12:{s:4:"Null";N;s:6:"String";s:5:"hello";s:7:"Php6Str";s:5:"world";s:3:"Int";i:123;s:4:"Bool";b:1;s:5:"False";b:0;s:5:"Float";d:12.34;s:6:"Object";O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}s:5:"ArrayValue";a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}s:9:"ObjectRef";R:9;s:8:"CSObject";O:3:"Bar":1:{s:3:"foo";s:54:"O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}";}s:6:"VarRef";R:3;}`
//	sCSObject = `O:3:"Bar":1:{s:3:"foo";s:54:"O:3:"Foo":2:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;}";}`
//	sVarRef   = `s:5:"hello";`
//)

type TestData struct {
	n  string             // Test Name
	f  phpcereal.TypeFlag // Type Flag
	e  bool               // escaped
	s  string             // Serialized String
	v  string             // Go Value
	t  string             // PHP Type
	r  []string           // Find/Replace strings
	i  bool               // Ignore string compare
	cr bool               // Count carriage Return
}

func (test TestData) IsCereal() bool {
	return phpcereal.IsCerealWithOpts(test.s, phpcereal.CerealOpts{
		Escaped: test.e,
		CountCR: test.cr,
	})
}

var testdata = []TestData{
	{
		n: "RevSlide Update Info",
		f: phpcereal.ObjectTypeFlag,
		e: false,
		s: `O:8:"stdClass":3:{s:7:"checked";i:1636591615;s:5:"basic";O:8:"stdClass":7:{s:4:"slug";s:9:"revslider";s:7:"version";s:5:"6.5.9";s:7:"package";s:102:"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e";s:6:"tested";s:5:"5.8.1";s:5:"icons";a:1:{s:2:"1x";s:62:"//updates.themepunch-ext-a.tools/revslider/logo.png?rev=6.2.23";}s:3:"url";s:33:"https://www.sliderrevolution.com/";s:11:"new_version";s:5:"6.5.9";}s:4:"full";O:8:"stdClass":19:{s:7:"banners";a:2:{s:3:"low";s:63:"//updates.themepunch-ext-a.tools/revslider/banner.png?rev=6.5.5";s:4:"high";s:63:"//updates.themepunch-ext-a.tools/revslider/banner.png?rev=6.5.5";}s:4:"name";s:17:"Slider Revolution";s:4:"slug";s:9:"revslider";s:6:"stable";s:5:"4.2.0";s:7:"version";s:5:"6.5.9";s:6:"tested";s:5:"5.8.1";s:14:"upgrade_notice";a:0:{}s:11:"last_update";s:10:"2021-10-11";s:12:"last_updated";s:10:"2021-10-11";s:8:"requires";s:3:"3.6";s:6:"author";s:51:"<a href="https://www.themepunch.com">ThemePunch</a>";s:7:"package";s:102:"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e";s:13:"download_link";s:102:"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e";s:9:"file_name";s:13:"revslider.zip";s:15:"active_installs";i:8000100;s:8:"homepage";s:33:"https://www.sliderrevolution.com/";s:8:"sections";a:3:{s:11:"description";s:497:"<h4>Slider Revolution WordPress Builder Plugin</h4>\n	<p>Slider Revolution 6 is a new way to build rich & dynamic content for your websites. With our powerful visual editor, you can create modern designs in no time, and with no coding experience required.</p><p>Create Sliders & Carousels, Hero Headers, Content Modules, Full Websites, Dynamic Solutions and Special FX with our amazing Add-Ons.</p>\n	<p>200+ templates are included in our online library. Cutting edge designs. Easily customized.</p>";s:9:"changelog";s:151:"<p>For Slider Revolution's changelog, please visit <a href="https://www.themepunch.com/slider-revolution/changelog/" target="_blank">this</a> site!</p>";s:3:"faq";s:2190:"<div class="tp-faq-content"><div class="tp-faq-column tp-faq-recent"><h4>Recent Solutions</h4><ul class="tp-faq-recent-content ready"><li><a href="https://www.youtube.com/watch?v=sCcnw5bZqYY&amp;list=PLSCdqDWVMJPPXEuOEqYEQMAsp0vAYw52_&amp;index=2&amp;t=111s" target="_blank" title="Video Tutorials">Video Tutorials</a></li><li><a href="https://www.themepunch.com/faq/responsive-content/" target="_blank" title="Responsive Content Setup">Responsive Content Setup</a></li><li><a href="https://www.themepunch.com/faq/video-content-mobile/" target="_blank" title="Video Content &amp; Mobile Considerations">Video Content &amp; Mobile Considerations</a></li><li><a href="https://www.themepunch.com/faq/how-to-change-the-timing-of-slides/" target="_blank" title="How to change the timing of Slides">How to change the timing of Slides</a></li><li><a href="https://www.themepunch.com/faq/mouse-hovers-for-layer-content/" target="_blank" title="Mouse Hovers for Layer Content"> Mouse Hovers for Layer Content</a></li></ul></div><div class="tp-faq-column tp-faq-popular"><h4>Popular Solutions</h4><ul class="tp-faq-popular-content ready"><li><a href="https://www.themepunch.com/faq/after-updating-make-sure-to-clear-all-caches/" target="_blank" title="After updating â Make sure to clear all caches">After updating â Make sure to clear all caches</a></li><li><a href="https://www.themepunch.com/faq/purchase-code-registration-faqs/" target="_blank" title="Purchase Code Registration FAQâs">Purchase Code Registration FAQâs</a></li><li><a href="https://www.themepunch.com/faq/ideal-image-size/" target="_blank" title="Ideal Image Size">Ideal Image Size</a></li><li><a href="https://www.themepunch.com/faq/add-links-to-slides-and-layers/" target="_blank" title="How to Hyperlink Slides and Layers">How to Hyperlink Slides and Layers</a></li><li><a href="https://www.themepunch.com/faq/where-to-find-the-purchase-code/" target="_blank" title="Where to find the Purchase Code">Where to find the Purchase Code</a></li></ul></div><div style="clear: both"></div><p><a class="button button-primary" href="https://themepunch.com/support-center"><strong>See All Faq's</strong></a></p></div>";}s:3:"url";s:33:"https://www.sliderrevolution.com/";s:8:"external";i:1;}}`,
		t: "stdClass",
		i: true,
	},
	{
		n: "Slider Revolution WordPress Builder Plugin Description",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:497:\"<h4>Slider Revolution WordPress Builder Plugin</h4>\n	<p>Slider Revolution 6 is a new way to build rich & dynamic content for your websites. With our powerful visual editor, you can create modern designs in no time, and with no coding experience required.</p><p>Create Sliders & Carousels, Hero Headers, Content Modules, Full Websites, Dynamic Solutions and Special FX with our amazing Add-Ons.</p>\n	<p>200+ templates are included in our online library. Cutting edge designs. Easily customized.</p>\";`,
		t: "string",
		i: true,
	},
	{
		n: "Object Properties of enquoted string with embedded quotes",
		f: phpcereal.ObjectTypeFlag,
		e: true,
		s: `O:8:\"stdClass\":1:{s:4:\"full\";O:8:\"stdClass\":1:{s:6:\"author\";s:51:\"<a href=\"https://www.themepunch.com\">ThemePunch</a>\";}}`,
		t: "stdClass",
		v: `stdClass{full:stdClass{author:\"<a href=\"https://www.themepunch.com\">ThemePunch</a>\"}}`,
	},
	{
		n: "Empty String Property",
		f: phpcereal.ObjectTypeFlag,
		t: "stdClass",
		e: true,
		s: `O:8:\"stdClass\":1:{s:3:\"foo\";s:0:\"\";}`,
		v: `stdClass{foo:\"\"}`,
	},
	{
		n: "Escaped Custom Object",
		f: phpcereal.CustomObjectTypeFlag,
		t: "WPSEO_Sitemap_Cache_Data",
		e: true,
		s: `C:24:\"WPSEO_Sitemap_Cache_Data\":583:{` +
			`a:2:{s:6:\"status\";s:2:\"ok\";s:3:\"xml\";s:536:\"<urlset xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" ` +
			`xmlns:image=\"http://www.google.com/schemas/sitemap-image/1.1\" xsi:schemaLocation=\"http://www.sitemaps.org/schemas/sitemap/0.9 ` +
			`http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd http://www.google.com/schemas/sitemap-image/1.1 ` +
			`http://www.google.com/schemas/sitemap-image/1.1/sitemap-image.xsd\" xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n	<url>` +
			`\n		<loc>https://www.examples.com/category/news/</loc>\n		<lastmod>2021-09-28T23:11:41+00:00</lastmod>\n	</url>\n</urlset>\";}` +
			`}`,
		v: `WPSEO_Sitemap_Cache_Data[` +
			`\"status\"=>\"ok\",` +
			`\"xml\"=>\"<urlset xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:image=\"http://www.google.com/schemas/sitemap-image/1.1\" ` +
			`xsi:schemaLocation=\"http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd ` +
			`http://www.google.com/schemas/sitemap-image/1.1 http://www.google.com/schemas/sitemap-image/1.1/sitemap-image.xsd\" ` +
			`xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n	<url>\n		<loc>https://www.examples.com/category/news/</loc>\n		` +
			`<lastmod>2021-09-28T23:11:41+00:00</lastmod>\n	</url>\n</urlset>\"` +
			`]`,
	},
	{
		n: "Embedded Escaped Control Characters",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:9:\"x\r\n	y\r\n	z\";`,
		t: "string",
		v: `\"x\r\n	y\r\n	z\"`,
		cr: true,
	},
	{
		n: "RevSlider Description Info Description",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:497:\"<h4>Slider Revolution WordPress Builder Plugin</h4>\n	<p>Slider Revolution 6 is a new way to build rich & dynamic content for your websites. With our powerful visual editor, you can create modern designs in no time, and with no coding experience required.</p><p>Create Sliders & Carousels, Hero Headers, Content Modules, Full Websites, Dynamic Solutions and Special FX with our amazing Add-Ons.</p>\n	<p>200+ templates are included in our online library. Cutting edge designs. Easily customized.</p>\";`,
		t: "string",
		i: true,
	},
	{
		n: "RevSlider Update Info Description with Slash R",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:497:\"<h4>Slider Revolution WordPress Builder Plugin</h4>\r\n	<p>Slider Revolution 6 is a new way to build rich & dynamic content for your websites. With our powerful visual editor, you can create modern designs in no time, and with no coding experience required.</p><p>Create Sliders & Carousels, Hero Headers, Content Modules, Full Websites, Dynamic Solutions and Special FX with our amazing Add-Ons.</p>\r\n	<p>200+ templates are included in our online library. Cutting edge designs. Easily customized.</p>\";`,
		t:  "string",
		i:  true,
		cr: false,
	},
	{
		n: "RevSlider Update Info",
		f: phpcereal.ObjectTypeFlag,
		e: true,
		s: `O:8:\"stdClass\":3:{s:7:\"checked\";i:1636591615;s:5:\"basic\";O:8:\"stdClass\":7:{s:4:\"slug\";s:9:\"revslider\";s:7:\"version\";s:5:\"6.5.9\";s:7:\"package\";s:102:\"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e\";s:6:\"tested\";s:5:\"5.8.1\";s:5:\"icons\";a:1:{s:2:\"1x\";s:62:\"//updates.themepunch-ext-a.tools/revslider/logo.png?rev=6.2.23\";}s:3:\"url\";s:33:\"https://www.sliderrevolution.com/\";s:11:\"new_version\";s:5:\"6.5.9\";}s:4:\"full\";O:8:\"stdClass\":19:{s:7:\"banners\";a:2:{s:3:\"low\";s:63:\"//updates.themepunch-ext-a.tools/revslider/banner.png?rev=6.5.5\";s:4:\"high\";s:63:\"//updates.themepunch-ext-a.tools/revslider/banner.png?rev=6.5.5\";}s:4:\"name\";s:17:\"Slider Revolution\";s:4:\"slug\";s:9:\"revslider\";s:6:\"stable\";s:5:\"4.2.0\";s:7:\"version\";s:5:\"6.5.9\";s:6:\"tested\";s:5:\"5.8.1\";s:14:\"upgrade_notice\";a:0:{}s:11:\"last_update\";s:10:\"2021-10-11\";s:12:\"last_updated\";s:10:\"2021-10-11\";s:8:\"requires\";s:3:\"3.6\";s:6:\"author\";s:51:\"<a href=\"https://www.themepunch.com\">ThemePunch</a>\";s:7:\"package\";s:102:\"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e\";s:13:\"download_link\";s:102:\"http://updates.themepunch-ext-a.tools/revslider/download.php?code=06db9573-b0ac-49e6-ac70-c2f21893446e\";s:9:\"file_name\";s:13:\"revslider.zip\";s:15:\"active_installs\";i:8000100;s:8:\"homepage\";s:33:\"https://www.sliderrevolution.com/\";s:8:\"sections\";a:3:{s:11:\"description\";s:497:\"<h4>Slider Revolution WordPress Builder Plugin</h4>\r\n	<p>Slider Revolution 6 is a new way to build rich & dynamic content for your websites. With our powerful visual editor, you can create modern designs in no time, and with no coding experience required.</p><p>Create Sliders & Carousels, Hero Headers, Content Modules, Full Websites, Dynamic Solutions and Special FX with our amazing Add-Ons.</p>\r\n	<p>200+ templates are included in our online library. Cutting edge designs. Easily customized.</p>\";s:9:\"changelog\";s:151:\"<p>For Slider Revolution\'s changelog, please visit <a href=\"https://www.themepunch.com/slider-revolution/changelog/\" target=\"_blank\">this</a> site!</p>\";s:3:\"faq\";s:2190:\"<div class=\"tp-faq-content\"><div class=\"tp-faq-column tp-faq-recent\"><h4>Recent Solutions</h4><ul class=\"tp-faq-recent-content ready\"><li><a href=\"https://www.youtube.com/watch?v=sCcnw5bZqYY&amp;list=PLSCdqDWVMJPPXEuOEqYEQMAsp0vAYw52_&amp;index=2&amp;t=111s\" target=\"_blank\" title=\"Video Tutorials\">Video Tutorials</a></li><li><a href=\"https://www.themepunch.com/faq/responsive-content/\" target=\"_blank\" title=\"Responsive Content Setup\">Responsive Content Setup</a></li><li><a href=\"https://www.themepunch.com/faq/video-content-mobile/\" target=\"_blank\" title=\"Video Content &amp; Mobile Considerations\">Video Content &amp; Mobile Considerations</a></li><li><a href=\"https://www.themepunch.com/faq/how-to-change-the-timing-of-slides/\" target=\"_blank\" title=\"How to change the timing of Slides\">How to change the timing of Slides</a></li><li><a href=\"https://www.themepunch.com/faq/mouse-hovers-for-layer-content/\" target=\"_blank\" title=\"Mouse Hovers for Layer Content\"> Mouse Hovers for Layer Content</a></li></ul></div><div class=\"tp-faq-column tp-faq-popular\"><h4>Popular Solutions</h4><ul class=\"tp-faq-popular-content ready\"><li><a href=\"https://www.themepunch.com/faq/after-updating-make-sure-to-clear-all-caches/\" target=\"_blank\" title=\"After updating â Make sure to clear all caches\">After updating â Make sure to clear all caches</a></li><li><a href=\"https://www.themepunch.com/faq/purchase-code-registration-faqs/\" target=\"_blank\" title=\"Purchase Code Registration FAQâs\">Purchase Code Registration FAQâs</a></li><li><a href=\"https://www.themepunch.com/faq/ideal-image-size/\" target=\"_blank\" title=\"Ideal Image Size\">Ideal Image Size</a></li><li><a href=\"https://www.themepunch.com/faq/add-links-to-slides-and-layers/\" target=\"_blank\" title=\"How to Hyperlink Slides and Layers\">How to Hyperlink Slides and Layers</a></li><li><a href=\"https://www.themepunch.com/faq/where-to-find-the-purchase-code/\" target=\"_blank\" title=\"Where to find the Purchase Code\">Where to find the Purchase Code</a></li></ul></div><div style=\"clear: both\"></div><p><a class=\"button button-primary\" href=\"https://themepunch.com/support-center\"><strong>See All Faq\'s</strong></a></p></div>\";}s:3:\"url\";s:33:\"https://www.sliderrevolution.com/\";s:8:\"external\";i:1;}}`,
		t:  "array",
		i:  true,
		cr: false,
	},
	{
		n: "Embedded carriage return as a single character",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:1:\"` + string('\n') + `\";`,
		t: "string",
		v: `\"` + string('\n') + `\"`,
	},
	{
		n: "Non-ascii chars in string",
		f: phpcereal.StringTypeFlag,
		e: false,
		s: `s:14:"Foo â Bar";`,
		t: "string",
		v: `"Foo â Bar"`,
	},
	{
		n: "Non-ascii chars in array element",
		f: phpcereal.ArrayTypeFlag,
		e: false,
		s: `a:1:{s:3:"faq";s:218:"<a href="https://www.themepunch.com/faq/after-updating-make-sure-to-clear-all-caches/" target="_blank" title="After updating â Make sure to clear all caches">After updating â Make sure to clear all caches</a>";}`,
		t: "array",
		v: `["faq"=>"<a href="https://www.themepunch.com/faq/after-updating-make-sure-to-clear-all-caches/" target="_blank" title="After updating â Make sure to clear all caches">After updating â Make sure to clear all caches</a>"]`,
	},
	{
		n: "String with embedded and unescaped control characters",
		f: phpcereal.StringTypeFlag,
		e: false,
		s: `s:11:"x` + string('\n') + string('\t') + `<p>x</p>";`,
		t: "string",
		v: `"x` + string('\n') + string('\t') + `<p>x</p>"`,
	},
	{
		n: "Enquoted string with embedded quotes",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:14:\"<a x=\"y\">z</a>\";`,
		t: "string",
		v: `\"<a x=\"y\">z</a>\"`,
	},
	{
		n: "Rewrite Rules",
		f: phpcereal.ArrayTypeFlag,
		e: true,
		s: `a:93:{s:11:\"^wp-json/?$\";s:22:\"index.php?rest_route=/\";s:14:\"^wp-json/(.*)?\";s:33:\"index.php?rest_route=/$matches[1]\";s:21:\"^index.php/wp-json/?$\";s:22:\"index.php?rest_route=/\";s:24:\"^index.php/wp-json/(.*)?\";s:33:\"index.php?rest_route=/$matches[1]\";s:17:\"^wp-sitemap\\.xml$\";s:23:\"index.php?sitemap=index\";s:17:\"^wp-sitemap\\.xsl$\";s:36:\"index.php?sitemap-stylesheet=sitemap\";s:23:\"^wp-sitemap-index\\.xsl$\";s:34:\"index.php?sitemap-stylesheet=index\";s:48:\"^wp-sitemap-([a-z]+?)-([a-z\\d_-]+?)-(\\d+?)\\.xml$\";s:75:\"index.php?sitemap=$matches[1]&sitemap-subtype=$matches[2]&paged=$matches[3]\";s:34:\"^wp-sitemap-([a-z]+?)-(\\d+?)\\.xml$\";s:47:\"index.php?sitemap=$matches[1]&paged=$matches[2]\";s:47:\"category/(.+?)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:52:\"index.php?category_name=$matches[1]&feed=$matches[2]\";s:42:\"category/(.+?)/(feed|rdf|rss|rss2|atom)/?$\";s:52:\"index.php?category_name=$matches[1]&feed=$matches[2]\";s:23:\"category/(.+?)/embed/?$\";s:46:\"index.php?category_name=$matches[1]&embed=true\";s:35:\"category/(.+?)/page/?([0-9]{1,})/?$\";s:53:\"index.php?category_name=$matches[1]&paged=$matches[2]\";s:17:\"category/(.+?)/?$\";s:35:\"index.php?category_name=$matches[1]\";s:44:\"tag/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:42:\"index.php?tag=$matches[1]&feed=$matches[2]\";s:39:\"tag/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:42:\"index.php?tag=$matches[1]&feed=$matches[2]\";s:20:\"tag/([^/]+)/embed/?$\";s:36:\"index.php?tag=$matches[1]&embed=true\";s:32:\"tag/([^/]+)/page/?([0-9]{1,})/?$\";s:43:\"index.php?tag=$matches[1]&paged=$matches[2]\";s:14:\"tag/([^/]+)/?$\";s:25:\"index.php?tag=$matches[1]\";s:45:\"type/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:50:\"index.php?post_format=$matches[1]&feed=$matches[2]\";s:40:\"type/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:50:\"index.php?post_format=$matches[1]&feed=$matches[2]\";s:21:\"type/([^/]+)/embed/?$\";s:44:\"index.php?post_format=$matches[1]&embed=true\";s:33:\"type/([^/]+)/page/?([0-9]{1,})/?$\";s:51:\"index.php?post_format=$matches[1]&paged=$matches[2]\";s:15:\"type/([^/]+)/?$\";s:33:\"index.php?post_format=$matches[1]\";s:12:\"robots\\.txt$\";s:18:\"index.php?robots=1\";s:13:\"favicon\\.ico$\";s:19:\"index.php?favicon=1\";s:48:\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2)\\.php$\";s:18:\"index.php?feed=old\";s:20:\".*wp-app\\.php(/.*)?$\";s:19:\"index.php?error=403\";s:18:\".*wp-register.php$\";s:23:\"index.php?register=true\";s:32:\"feed/(feed|rdf|rss|rss2|atom)/?$\";s:27:\"index.php?&feed=$matches[1]\";s:27:\"(feed|rdf|rss|rss2|atom)/?$\";s:27:\"index.php?&feed=$matches[1]\";s:8:\"embed/?$\";s:21:\"index.php?&embed=true\";s:20:\"page/?([0-9]{1,})/?$\";s:28:\"index.php?&paged=$matches[1]\";s:41:\"comments/feed/(feed|rdf|rss|rss2|atom)/?$\";s:42:\"index.php?&feed=$matches[1]&withcomments=1\";s:36:\"comments/(feed|rdf|rss|rss2|atom)/?$\";s:42:\"index.php?&feed=$matches[1]&withcomments=1\";s:17:\"comments/embed/?$\";s:21:\"index.php?&embed=true\";s:44:\"search/(.+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:40:\"index.php?s=$matches[1]&feed=$matches[2]\";s:39:\"search/(.+)/(feed|rdf|rss|rss2|atom)/?$\";s:40:\"index.php?s=$matches[1]&feed=$matches[2]\";s:20:\"search/(.+)/embed/?$\";s:34:\"index.php?s=$matches[1]&embed=true\";s:32:\"search/(.+)/page/?([0-9]{1,})/?$\";s:41:\"index.php?s=$matches[1]&paged=$matches[2]\";s:14:\"search/(.+)/?$\";s:23:\"index.php?s=$matches[1]\";s:47:\"author/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:50:\"index.php?author_name=$matches[1]&feed=$matches[2]\";s:42:\"author/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:50:\"index.php?author_name=$matches[1]&feed=$matches[2]\";s:23:\"author/([^/]+)/embed/?$\";s:44:\"index.php?author_name=$matches[1]&embed=true\";s:35:\"author/([^/]+)/page/?([0-9]{1,})/?$\";s:51:\"index.php?author_name=$matches[1]&paged=$matches[2]\";s:17:\"author/([^/]+)/?$\";s:33:\"index.php?author_name=$matches[1]\";s:69:\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/feed/(feed|rdf|rss|rss2|atom)/?$\";s:80:\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&feed=$matches[4]\";s:64:\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/(feed|rdf|rss|rss2|atom)/?$\";s:80:\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&feed=$matches[4]\";s:45:\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/embed/?$\";s:74:\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&embed=true\";s:57:\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/page/?([0-9]{1,})/?$\";s:81:\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&paged=$matches[4]\";s:39:\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/?$\";s:63:\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]\";s:56:\"([0-9]{4})/([0-9]{1,2})/feed/(feed|rdf|rss|rss2|atom)/?$\";s:64:\"index.php?year=$matches[1]&monthnum=$matches[2]&feed=$matches[3]\";s:51:\"([0-9]{4})/([0-9]{1,2})/(feed|rdf|rss|rss2|atom)/?$\";s:64:\"index.php?year=$matches[1]&monthnum=$matches[2]&feed=$matches[3]\";s:32:\"([0-9]{4})/([0-9]{1,2})/embed/?$\";s:58:\"index.php?year=$matches[1]&monthnum=$matches[2]&embed=true\";s:44:\"([0-9]{4})/([0-9]{1,2})/page/?([0-9]{1,})/?$\";s:65:\"index.php?year=$matches[1]&monthnum=$matches[2]&paged=$matches[3]\";s:26:\"([0-9]{4})/([0-9]{1,2})/?$\";s:47:\"index.php?year=$matches[1]&monthnum=$matches[2]\";s:43:\"([0-9]{4})/feed/(feed|rdf|rss|rss2|atom)/?$\";s:43:\"index.php?year=$matches[1]&feed=$matches[2]\";s:38:\"([0-9]{4})/(feed|rdf|rss|rss2|atom)/?$\";s:43:\"index.php?year=$matches[1]&feed=$matches[2]\";s:19:\"([0-9]{4})/embed/?$\";s:37:\"index.php?year=$matches[1]&embed=true\";s:31:\"([0-9]{4})/page/?([0-9]{1,})/?$\";s:44:\"index.php?year=$matches[1]&paged=$matches[2]\";s:13:\"([0-9]{4})/?$\";s:26:\"index.php?year=$matches[1]\";s:27:\".?.+?/attachment/([^/]+)/?$\";s:32:\"index.php?attachment=$matches[1]\";s:37:\".?.+?/attachment/([^/]+)/trackback/?$\";s:37:\"index.php?attachment=$matches[1]&tb=1\";s:57:\".?.+?/attachment/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:52:\".?.+?/attachment/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:52:\".?.+?/attachment/([^/]+)/comment-page-([0-9]{1,})/?$\";s:50:\"index.php?attachment=$matches[1]&cpage=$matches[2]\";s:33:\".?.+?/attachment/([^/]+)/embed/?$\";s:43:\"index.php?attachment=$matches[1]&embed=true\";s:16:\"(.?.+?)/embed/?$\";s:41:\"index.php?pagename=$matches[1]&embed=true\";s:20:\"(.?.+?)/trackback/?$\";s:35:\"index.php?pagename=$matches[1]&tb=1\";s:40:\"(.?.+?)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:47:\"index.php?pagename=$matches[1]&feed=$matches[2]\";s:35:\"(.?.+?)/(feed|rdf|rss|rss2|atom)/?$\";s:47:\"index.php?pagename=$matches[1]&feed=$matches[2]\";s:28:\"(.?.+?)/page/?([0-9]{1,})/?$\";s:48:\"index.php?pagename=$matches[1]&paged=$matches[2]\";s:35:\"(.?.+?)/comment-page-([0-9]{1,})/?$\";s:48:\"index.php?pagename=$matches[1]&cpage=$matches[2]\";s:24:\"(.?.+?)(?:/([0-9]+))?/?$\";s:47:\"index.php?pagename=$matches[1]&page=$matches[2]\";s:27:\"[^/]+/attachment/([^/]+)/?$\";s:32:\"index.php?attachment=$matches[1]\";s:37:\"[^/]+/attachment/([^/]+)/trackback/?$\";s:37:\"index.php?attachment=$matches[1]&tb=1\";s:57:\"[^/]+/attachment/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:52:\"[^/]+/attachment/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:52:\"[^/]+/attachment/([^/]+)/comment-page-([0-9]{1,})/?$\";s:50:\"index.php?attachment=$matches[1]&cpage=$matches[2]\";s:33:\"[^/]+/attachment/([^/]+)/embed/?$\";s:43:\"index.php?attachment=$matches[1]&embed=true\";s:16:\"([^/]+)/embed/?$\";s:37:\"index.php?name=$matches[1]&embed=true\";s:20:\"([^/]+)/trackback/?$\";s:31:\"index.php?name=$matches[1]&tb=1\";s:40:\"([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:43:\"index.php?name=$matches[1]&feed=$matches[2]\";s:35:\"([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:43:\"index.php?name=$matches[1]&feed=$matches[2]\";s:28:\"([^/]+)/page/?([0-9]{1,})/?$\";s:44:\"index.php?name=$matches[1]&paged=$matches[2]\";s:35:\"([^/]+)/comment-page-([0-9]{1,})/?$\";s:44:\"index.php?name=$matches[1]&cpage=$matches[2]\";s:24:\"([^/]+)(?:/([0-9]+))?/?$\";s:43:\"index.php?name=$matches[1]&page=$matches[2]\";s:16:\"[^/]+/([^/]+)/?$\";s:32:\"index.php?attachment=$matches[1]\";s:26:\"[^/]+/([^/]+)/trackback/?$\";s:37:\"index.php?attachment=$matches[1]&tb=1\";s:46:\"[^/]+/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:41:\"[^/]+/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\";s:49:\"index.php?attachment=$matches[1]&feed=$matches[2]\";s:41:\"[^/]+/([^/]+)/comment-page-([0-9]{1,})/?$\";s:50:\"index.php?attachment=$matches[1]&cpage=$matches[2]\";s:22:\"[^/]+/([^/]+)/embed/?$\";s:43:\"index.php?attachment=$matches[1]&embed=true\";}`,
		t: "array",
		v: `[` +
			`\"^wp-json/?$\"=>\"index.php?rest_route=/\",` +
			`\"^wp-json/(.*)?\"=>\"index.php?rest_route=/$matches[1]\",` +
			`\"^index.php/wp-json/?$\"=>\"index.php?rest_route=/\",` +
			`\"^index.php/wp-json/(.*)?\"=>\"index.php?rest_route=/$matches[1]\",` +
			`\"^wp-sitemap\\.xml$\"=>\"index.php?sitemap=index\",` +
			`\"^wp-sitemap\\.xsl$\"=>\"index.php?sitemap-stylesheet=sitemap\",` +
			`\"^wp-sitemap-index\\.xsl$\"=>\"index.php?sitemap-stylesheet=index\",` +
			`\"^wp-sitemap-([a-z]+?)-([a-z\\d_-]+?)-(\\d+?)\\.xml$\"=>\"index.php?sitemap=$matches[1]&sitemap-subtype=$matches[2]&paged=$matches[3]\",` +
			`\"^wp-sitemap-([a-z]+?)-(\\d+?)\\.xml$\"=>\"index.php?sitemap=$matches[1]&paged=$matches[2]\",` +
			`\"category/(.+?)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?category_name=$matches[1]&feed=$matches[2]\",` +
			`\"category/(.+?)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?category_name=$matches[1]&feed=$matches[2]\",` +
			`\"category/(.+?)/embed/?$\"=>\"index.php?category_name=$matches[1]&embed=true\",` +
			`\"category/(.+?)/page/?([0-9]{1,})/?$\"=>\"index.php?category_name=$matches[1]&paged=$matches[2]\",` +
			`\"category/(.+?)/?$\"=>\"index.php?category_name=$matches[1]\",` +
			`\"tag/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?tag=$matches[1]&feed=$matches[2]\",` +
			`\"tag/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?tag=$matches[1]&feed=$matches[2]\",` +
			`\"tag/([^/]+)/embed/?$\"=>\"index.php?tag=$matches[1]&embed=true\",` +
			`\"tag/([^/]+)/page/?([0-9]{1,})/?$\"=>\"index.php?tag=$matches[1]&paged=$matches[2]\",` +
			`\"tag/([^/]+)/?$\"=>\"index.php?tag=$matches[1]\",` +
			`\"type/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?post_format=$matches[1]&feed=$matches[2]\",` +
			`\"type/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?post_format=$matches[1]&feed=$matches[2]\",` +
			`\"type/([^/]+)/embed/?$\"=>\"index.php?post_format=$matches[1]&embed=true\",` +
			`\"type/([^/]+)/page/?([0-9]{1,})/?$\"=>\"index.php?post_format=$matches[1]&paged=$matches[2]\",` +
			`\"type/([^/]+)/?$\"=>\"index.php?post_format=$matches[1]\",` +
			`\"robots\\.txt$\"=>\"index.php?robots=1\",` +
			`\"favicon\\.ico$\"=>\"index.php?favicon=1\",` +
			`\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2)\\.php$\"=>\"index.php?feed=old\",` +
			`\".*wp-app\\.php(/.*)?$\"=>\"index.php?error=403\",` +
			`\".*wp-register.php$\"=>\"index.php?register=true\",` +
			`\"feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?&feed=$matches[1]\",` +
			`\"(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?&feed=$matches[1]\",` +
			`\"embed/?$\"=>\"index.php?&embed=true\",` +
			`\"page/?([0-9]{1,})/?$\"=>\"index.php?&paged=$matches[1]\",` +
			`\"comments/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?&feed=$matches[1]&withcomments=1\",` +
			`\"comments/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?&feed=$matches[1]&withcomments=1\",` +
			`\"comments/embed/?$\"=>\"index.php?&embed=true\",` +
			`\"search/(.+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?s=$matches[1]&feed=$matches[2]\",` +
			`\"search/(.+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?s=$matches[1]&feed=$matches[2]\",` +
			`\"search/(.+)/embed/?$\"=>\"index.php?s=$matches[1]&embed=true\",` +
			`\"search/(.+)/page/?([0-9]{1,})/?$\"=>\"index.php?s=$matches[1]&paged=$matches[2]\",` +
			`\"search/(.+)/?$\"=>\"index.php?s=$matches[1]\",` +
			`\"author/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?author_name=$matches[1]&feed=$matches[2]\",` +
			`\"author/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?author_name=$matches[1]&feed=$matches[2]\",` +
			`\"author/([^/]+)/embed/?$\"=>\"index.php?author_name=$matches[1]&embed=true\",` +
			`\"author/([^/]+)/page/?([0-9]{1,})/?$\"=>\"index.php?author_name=$matches[1]&paged=$matches[2]\",` +
			`\"author/([^/]+)/?$\"=>\"index.php?author_name=$matches[1]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&feed=$matches[4]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&feed=$matches[4]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/embed/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&embed=true\",` +
			`\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/page/?([0-9]{1,})/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]&paged=$matches[4]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&day=$matches[3]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&feed=$matches[3]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&feed=$matches[3]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/embed/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&embed=true\",` +
			`\"([0-9]{4})/([0-9]{1,2})/page/?([0-9]{1,})/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]&paged=$matches[3]\",` +
			`\"([0-9]{4})/([0-9]{1,2})/?$\"=>\"index.php?year=$matches[1]&monthnum=$matches[2]\",` +
			`\"([0-9]{4})/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&feed=$matches[2]\",` +
			`\"([0-9]{4})/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?year=$matches[1]&feed=$matches[2]\",` +
			`\"([0-9]{4})/embed/?$\"=>\"index.php?year=$matches[1]&embed=true\",` +
			`\"([0-9]{4})/page/?([0-9]{1,})/?$\"=>\"index.php?year=$matches[1]&paged=$matches[2]\",` +
			`\"([0-9]{4})/?$\"=>\"index.php?year=$matches[1]\",` +
			`\".?.+?/attachment/([^/]+)/?$\"=>\"index.php?attachment=$matches[1]\",` +
			`\".?.+?/attachment/([^/]+)/trackback/?$\"=>\"index.php?attachment=$matches[1]&tb=1\",` +
			`\".?.+?/attachment/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\".?.+?/attachment/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\".?.+?/attachment/([^/]+)/comment-page-([0-9]{1,})/?$\"=>\"index.php?attachment=$matches[1]&cpage=$matches[2]\",` +
			`\".?.+?/attachment/([^/]+)/embed/?$\"=>\"index.php?attachment=$matches[1]&embed=true\",` +
			`\"(.?.+?)/embed/?$\"=>\"index.php?pagename=$matches[1]&embed=true\",` +
			`\"(.?.+?)/trackback/?$\"=>\"index.php?pagename=$matches[1]&tb=1\",` +
			`\"(.?.+?)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?pagename=$matches[1]&feed=$matches[2]\",` +
			`\"(.?.+?)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?pagename=$matches[1]&feed=$matches[2]\",` +
			`\"(.?.+?)/page/?([0-9]{1,})/?$\"=>\"index.php?pagename=$matches[1]&paged=$matches[2]\",` +
			`\"(.?.+?)/comment-page-([0-9]{1,})/?$\"=>\"index.php?pagename=$matches[1]&cpage=$matches[2]\",` +
			`\"(.?.+?)(?:/([0-9]+))?/?$\"=>\"index.php?pagename=$matches[1]&page=$matches[2]\",` +
			`\"[^/]+/attachment/([^/]+)/?$\"=>\"index.php?attachment=$matches[1]\",` +
			`\"[^/]+/attachment/([^/]+)/trackback/?$\"=>\"index.php?attachment=$matches[1]&tb=1\",` +
			`\"[^/]+/attachment/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\"[^/]+/attachment/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\"[^/]+/attachment/([^/]+)/comment-page-([0-9]{1,})/?$\"=>\"index.php?attachment=$matches[1]&cpage=$matches[2]\",` +
			`\"[^/]+/attachment/([^/]+)/embed/?$\"=>\"index.php?attachment=$matches[1]&embed=true\",` +
			`\"([^/]+)/embed/?$\"=>\"index.php?name=$matches[1]&embed=true\",` +
			`\"([^/]+)/trackback/?$\"=>\"index.php?name=$matches[1]&tb=1\",` +
			`\"([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?name=$matches[1]&feed=$matches[2]\",` +
			`\"([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?name=$matches[1]&feed=$matches[2]\",` +
			`\"([^/]+)/page/?([0-9]{1,})/?$\"=>\"index.php?name=$matches[1]&paged=$matches[2]\",` +
			`\"([^/]+)/comment-page-([0-9]{1,})/?$\"=>\"index.php?name=$matches[1]&cpage=$matches[2]\",` +
			`\"([^/]+)(?:/([0-9]+))?/?$\"=>\"index.php?name=$matches[1]&page=$matches[2]\",` +
			`\"[^/]+/([^/]+)/?$\"=>\"index.php?attachment=$matches[1]\",` +
			`\"[^/]+/([^/]+)/trackback/?$\"=>\"index.php?attachment=$matches[1]&tb=1\",` +
			`\"[^/]+/([^/]+)/feed/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\"[^/]+/([^/]+)/(feed|rdf|rss|rss2|atom)/?$\"=>\"index.php?attachment=$matches[1]&feed=$matches[2]\",` +
			`\"[^/]+/([^/]+)/comment-page-([0-9]{1,})/?$\"=>\"index.php?attachment=$matches[1]&cpage=$matches[2]\",` +
			`\"[^/]+/([^/]+)/embed/?$\"=>\"index.php?attachment=$matches[1]&embed=true\"` +
			`]`,
	},
	{
		n: "Escaped Complex Array Value",
		f: phpcereal.ArrayTypeFlag,
		t: "array",
		e: true,
		s: `a:24:{` +
			`i:0;s:34:\"advanced-custom-fields-pro/acf.php\";` +
			`i:1;s:55:\"advanced-post-types-order/advanced-post-types-order.php\";` +
			`i:2;s:39:\"aryo-activity-log/aryo-activity-log.php\";` +
			`i:3;s:28:\"category-posts/cat-posts.php\";` +
			`i:4;s:33:\"classic-editor/classic-editor.php\";` +
			`i:5;s:49:\"constant-contact-forms/constant-contact-forms.php\";` +
			`i:6;s:35:\"disable-xml-rpc/disable-xml-rpc.php\";` +
			`i:7;s:61:\"divi_extended_column_layouts/divi_extended_column_layouts.php\";` +
			`i:8;s:32:\"duplicate-page/duplicatepage.php\";` +
			`i:9;s:49:\"elegant-themes-updater/elegant-themes-updater.php\";` +
			`i:10;s:45:\"enable-media-replace/enable-media-replace.php\";` +
			`i:11;s:53:\"enhanced-media-library-pro/enhanced-media-library.php\";` +
			`i:12;s:21:\"include-me/plugin.php\";` +
			`i:13;s:30:\"interactive-world-maps/map.php\";` +
			`i:14;s:63:\"limit-login-attempts-reloaded/limit-login-attempts-reloaded.php\";` +
			`i:15;s:25:\"menu-image/menu-image.php\";` +
			`i:16;s:33:\"posts-to-posts/posts-to-posts.php\";` +
			`i:17;s:19:\"monarch/monarch.php\";` +
			`i:18;s:47:\"regenerate-thumbnails/regenerate-thumbnails.php\";` +
			`i:19;s:27:\"redirection/redirection.php\";` +
			`i:20;s:23:\"revslider/revslider.php\";` +
			`i:21;s:45:\"taxonomy-terms-order/taxonomy-terms-order.php\";` +
			`i:22;s:24:\"wordpress-seo/wp-seo.php\";` +
			`i:23;s:29:\"wp-shopify-pro/wp-shopify.php\";` +
			`}`,
		v: `[` +
			`0=>\"advanced-custom-fields-pro/acf.php\",` +
			`1=>\"advanced-post-types-order/advanced-post-types-order.php\",` +
			`2=>\"aryo-activity-log/aryo-activity-log.php\",` +
			`3=>\"category-posts/cat-posts.php\",` +
			`4=>\"classic-editor/classic-editor.php\",` +
			`5=>\"constant-contact-forms/constant-contact-forms.php\",` +
			`6=>\"disable-xml-rpc/disable-xml-rpc.php\",` +
			`7=>\"divi_extended_column_layouts/divi_extended_column_layouts.php\",` +
			`8=>\"duplicate-page/duplicatepage.php\",` +
			`9=>\"elegant-themes-updater/elegant-themes-updater.php\",` +
			`10=>\"enable-media-replace/enable-media-replace.php\",` +
			`11=>\"enhanced-media-library-pro/enhanced-media-library.php\",` +
			`12=>\"include-me/plugin.php\",` +
			`13=>\"interactive-world-maps/map.php\",` +
			`14=>\"limit-login-attempts-reloaded/limit-login-attempts-reloaded.php\",` +
			`15=>\"menu-image/menu-image.php\",` +
			`16=>\"posts-to-posts/posts-to-posts.php\",` +
			`17=>\"monarch/monarch.php\",` +
			`18=>\"regenerate-thumbnails/regenerate-thumbnails.php\",` +
			`19=>\"redirection/redirection.php\",` +
			`20=>\"revslider/revslider.php\",` +
			`21=>\"taxonomy-terms-order/taxonomy-terms-order.php\",` +
			`22=>\"wordpress-seo/wp-seo.php\",` +
			`23=>\"wp-shopify-pro/wp-shopify.php\"` +
			`]`,
	},
	{
		n: "Non-escaped Embedded Slash N",
		f: phpcereal.StringTypeFlag,
		e: false,
		s: `s:5:"Test\n";`,
		t: "string",
		v: `"Test\n"`, // Legend: ~private, ^protected
	},
	{
		n: "Escaped string world",
		f: phpcereal.PHP6StringTypeFlag,
		e: true,
		s: `S:5:\"world\";`,
		v: `\"world\"`,
		t: "6string",
	},
	{
		n: "Regex-escaped characters",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:17:\"^wp-sitemap\\.xml$\";`,
		t: "string",
		v: `\"^wp-sitemap\\.xml$\"`, // Legend: ~private, ^protected
	},
	{
		n: "Object: Foo",
		f: phpcereal.ObjectTypeFlag,
		s: `O:3:"Foo":3:{s:3:"foo";s:3:"abc";s:8:"\0Foo\0bar";i:13;s:6:"\0*\0baz";b:1;}`,
		v: `Foo{foo:"abc",~bar:13,^baz:true}`, // Legend: ~private, ^protected
		t: "Foo",
	},
	{
		n: "UnescapedEmptyString",
		f: phpcereal.StringTypeFlag,
		t: "string",
		e: false,
		s: `s:0:"";`,
		v: `""`,
	},
	{
		n: "EscapedEmptyString",
		f: phpcereal.StringTypeFlag,
		t: "string",
		e: true,
		s: `s:0:\"\";`,
		v: `\"\"`,
	},
	{
		n: "EscapedEmptyStringAsArrayElement",
		f: phpcereal.ArrayTypeFlag,
		t: "array",
		e: true,
		s: `a:1:{i:0;s:0:\"\";}`,
		v: `[0=>\"\"]`,
	},
	{
		n: "Short Custom Object",
		f: phpcereal.CustomObjectTypeFlag,
		t: "Student",
		s: `C:7:"Student":27:{a:1:{s:4:"name";s:3:"Bob";}}`,
		v: `Student["name"=>"Bob"]`,
	},
	{
		n: "Custom Object",
		f: phpcereal.CustomObjectTypeFlag,
		t: "Student",
		s: `C:7:"Student":73:{a:4:{s:4:"name";s:3:"Bob";s:5:"class";i:10;s:4:"roll";i:3;s:3:"gpa";d:4;}}`,
		v: `Student["name"=>"Bob","class"=>10,"roll"=>3,"gpa"=>4.0]`,
	},
	{
		n: "Float:12.34",
		f: phpcereal.FloatTypeFlag,
		s: "d:12.34;",
		v: "12.34",
		t: "float",
	},
	{
		n: "ArrayValue of URLs",
		f: phpcereal.ArrayTypeFlag,
		s: `a:3:{i:0;s:40:"https://en.wiktionary.org/wiki/enquoting";i:1;s:41:"https://en.wiktionary.org/wiki/whiff#Verb";i:2;s:42:"https://en.wiktionary.org/wiki/tea#Spanish";}`,
		v: `[0=>"https://en.wikipedia.org/wiki/enquoting",1=>"https://en.wikipedia.org/wiki/whiff#Verb",2=>"https://en.wikipedia.org/wiki/tea#Spanish"]`,
		t: "array",
		r: []string{"wiktionary.org", "wikipedia.org"},
	},
	{
		n: "ArrayValue of Integers:[1,2,3]",
		f: phpcereal.ArrayTypeFlag,
		s: "a:3:{i:0;i:1;i:1;i:2;i:2;i:3;}",
		v: "[0=>1,1=>2,2=>3]",
		t: "array",
	},
	{
		n: "String:hello",
		f: phpcereal.StringTypeFlag,
		s: `s:5:"hello";`,
		v: `"herro"`,
		t: "string",
		r: []string{"ll", "rr"},
	},
	{
		n: "Non-escaped string world",
		f: phpcereal.PHP6StringTypeFlag,
		e: false,
		s: `S:5:"world";`,
		v: `"world"`,
		t: "6string",
	},
	{
		n: "NULL",
		f: phpcereal.NULLTypeFlag,
		s: `N;`,
		v: "NULL",
		t: "NULL",
	},
	{
		n: "Bool:true",
		f: phpcereal.BoolTypeFlag,
		s: "b:1;",
		v: "true",
		t: "bool",
	},
	{
		n: "Bool:false",
		f: phpcereal.BoolTypeFlag,
		s: "b:0;",
		v: "false",
		t: "bool",
	},
	{
		n: "Integer:123",
		f: phpcereal.IntTypeFlag,
		s: "i:123;",
		v: "123",
		t: "int",
	},
	{
		n: "Escaped Embedded Slash N",
		f: phpcereal.StringTypeFlag,
		e: true,
		s: `s:5:\"Test\n\";`,
		t: "string",
		v: `\"Test\n\"`, // Legend: ~private, ^protected
	},
	// ==========vvvv==PASSING==vvvv==========

	//==========^^^^==FAILING==^^^^==========
}

func TestParsing(t *testing.T) {
	var s string

	for _, test := range testdata {
		t.Run(test.n, func(t *testing.T) {
			if !test.IsCereal() {
				t.Errorf("failed to validate: %s", test.s)
			}
			sp := phpcereal.NewParser(test.s)
			sp.SetOpts(phpcereal.CerealOpts{
				Escaped: test.e,
				CountCR: test.cr,
			})
			root, err := sp.Parse()
			if err != nil {
				t.Error(err.Error())
				return
			}
			s = root.Serialized()
			assert.Equal(t, test.s, s)
			if assert.NotEmpty(t, s) {
				assert.Equal(t, test.f, phpcereal.TypeFlag(s[0]),
					"Type Flag not as expected: %q <> %q",
					test.f,
					phpcereal.TypeFlag(s[0]),
				)
			} else {
				return
			}
			if len(test.r) == 2 {
				if sr, ok := root.(phpcereal.StringReplacer); ok {
					sr.ReplaceString(test.r[0], test.r[1], -1)
				}
			}
			if test.i {
				return
			}
			rt := root.GetType()
			if assert.Equal(t, phpcereal.PHPType(test.t), rt) {
				assert.Equal(t, test.v, root.String())
			} else {
				t.Errorf("mismatch in types: %s <> %s", phpcereal.PHPType(test.t), root.GetType())
			}
		})
	}
}
