package ealing

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseAddresses(t *testing.T) {

	results := parseAddresses(addressesResponse)

	assert.NotNil(t, results)
	assert.NotEmpty(t, results)
	assert.Len(t, results, 58)


	assert.Equal(t, "12125514        ", results[0].UPRN)
	assert.Equal(t, "DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU", results[0].Address)

	assert.Equal(t, "12003001        ", results[57].UPRN)
	assert.Equal(t, "57, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU", results[57].Address)

}





var addressesResponse = `<!DOCTYPE html>
<!--[if lt IE 7]>       <html class="no-js lt-ie10 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>          <html class="no-js lt-ie10 lt-ie9 lt-ie8" lang="en"> <![endif]-->
<!--[if IE 8]>          <html class="no-js lt-ie10 lt-ie9" lang="en"> <![endif]-->
<!--[if IE 9]>          <html class="no-js lt-ie10" lang="en"> <![endif]-->
<!--[if gt IE 9]><!-->  <html class="no-js" lang="en"> <!--<![endif]-->
    <head>
      
        <meta http-equiv="x-ua-compatible" content="IE=edge,Chrome=1">
      
        <link rel="alternate" type="application/rss+xml" title="RSS" href="/rss/news">
        <link rel="apple-touch-icon" href="/site/images/apple-touch-icon.png">
        <link rel="search" type="application/opensearchdescription+xml" title="Ealing Council" href="/site/scripts/opensearch.php">
        <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
<link rel="stylesheet" type="text/css" href="/site/styles/generic/base.css">

<!--[if lt IE 9]>
    <link href="/site/dist/ealing-oldie.css" rel="stylesheet">
<![endif]-->
<!--[if gte IE 9]><!-->
    <link href="/site/dist/ealing.css" rel="stylesheet">
<!--<![endif]-->



        <script src="/site/libs/modernizr/modernizr.js"></script>
        <script src="/site/javascript/pseudoelementlineheightinrems.js"></script>
        <!--[if lt IE 9]>
        <script src="/site/javascript/selectivizr-min.js"></script>
        <![endif]-->
        <meta name="author" content="Ealing Council">
        <meta name="generator" content="http://www.jadu.net">
        <meta name="revisit-after" content="2 days">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <meta name="ROBOTS" content="INDEX,FOLLOW" >

	<link rel="stylesheet" type="text/css" href="/site/xfp/styles/generic/forms.css" media="screen" />

		<meta name="Keywords" content="Ealing Council" />
	<meta name="Description" content="Ealing Council - Report Problem" />
	<!-- HTML5 Template -->

	<!-- not HTML5 suppported -->
	<!-- Dublin Core Metadata -->
<meta name="dcterms.creator" content="Ealing Council" >
<meta name="dc.created" content="2012-09-14" >
<meta name="dc.modified" content="2012-09-14" >
<meta name="dc.valid" content="2012-09-14" >
<meta name="dc.description" content="Online form to check rubbish and recycling collection day" >
<meta name="dcterms.format" content="text/html" >
<meta name="dcterms.identifier" content="/site/custom_scripts/waste_collection/waste_collection.aspx" >
<meta name="dcterms.language" content="en" >
<meta name="dcterms.publisher" content="Ealing Council" >
<meta name="dcterms.rights" content="Copyright 2011 Ealing Council" >
<meta name="dc.title" content="Rubbish and recycling collection day" >
<meta name="dcterms.coverage" content="Hanwell, Ealing, Greenford, Perivale, Northolt, Acton, Southall, London, England, United Kingdom" >
<meta name="dcterms.subject" content="rubbish, waste, recycling" >


         <script type="text/javascript" src="/site/javascript/swfobject.js"></script>
<title>Rubbish and recycling collection day</title>


        <script>
            (function (html) {
                html.className = html.className.replace(/\bno-js\b/, '');
            })(document.getElementsByTagName('html')[0]);
        </script>
    </head>

    <body class="one-column"  id="xfp">

    <!-- googleoff: index -->
        <ul class="skip-links">
            <li>
                <a href="#content" rel="nofollow">Skip to content</a>
            </li>
        </ul>


        <header role="banner" class="site-header" id="top">
            <div class="container">
                <div class="site-header__logo-wrapper">
                    <a class="logo" href="/">
                        <span class="h1">Ealing Council</span>
                    </a>
                </div>
                

                            <ul class="site-header__links item-list item-list--inline" id="site-navigation">
                    <li>
                        <a href="/a_to_z/">A-Z</a>
                    </li>
                    <li>
                        <a href="/homepage/28/jobs">Jobs</a>
                    </li>
                    <li>
                        <a href="/homepage/182/do_it_online">Do it online</a>
                    </li>
                    <li>
                        <a href="/homepage/32/contact_us">Contact us</a>
                    </li>
                    <li>
                        <a href="/signin">Register/My Account</a>
                    </li>
                </ul>


                <div class="site-toggles">
                    <button class="search-toggle button button--standout" data-overlay="site-search">Search</button>
                    <button class="menu-toggle button" data-overlay="site-navigation">Menu</button>
                </div>
            </div>


 </header>

<div class="advert">
    <div>
        <script type="text/javascript">
            document.write('\x3Cscript type="text/javascript" src="//can.capacitygrid.com/code/ealinglbc/leaderboard/public">\x3C/script>');
    </script>
    </div>
</div>

<!-- googleon: index -->
            <main class="site-main" role="main">
                <h1 class="page-heading">
                    <span class="page-heading__text container">
                        Rubbish and recycling collection day</span>

                </h1>

<!-- googleoff: all -->



            <div class="container">




                <div class="site-content" id="content" tabindex="-1">
<!-- googleon: all -->


 <!-- ################ START OF WASTE COLLECTION FORM ############ -->

 <script type="text/javascript">
     function ShowHide(obj) {
         var parent = obj + '-parent';
         if ($(obj).style.display == 'none') {
             new Effect.SlideDown(obj, { duration: 0.5 });
             document.cookie = obj + "=1; path=/; domain=.ealing.gov.uk; secure";
             document.cookie = obj + "=1; path=/; domain=.ealing.gov.uk;";
             $(parent).addClassName('accordion-toggle-active');
         }
         else {
             new Effect.SlideUp(obj, { duration: 0.5 });
             document.cookie = obj + "=0; path=/; domain=.ealing.gov.uk; secure";
             document.cookie = obj + "=0; path=/; domain=.ealing.gov.uk;";
             $(parent).removeClassName('accordion-toggle-active');
         }
     }


     //when browser loads populate Address fields
     if (typeof window.onload == 'function') { detectCheckbox = window.onload; }
     window.onload = function () {
         getSelectedAddress();
         timedCollectionInfo();
     }

     //Timed Collections details are displayed above the schedule grid
     // if a property doesn't have timed collections, remove the div to
     // prevent a big ol' empty space being left there...'
     function timedCollectionInfo() {
         if (document.getElementById("hiddenTimedCollectionCode").value == '') {
             $("#TimedCollectionInfo").hide();
         }
     }


     function getSelectedAddress() {

         document.getElementById("hiddenUPRN").value = "no";
         document.getElementById("txtHouseNumber").value = "";
         document.getElementById("txtBuildingName").value = "";
         document.getElementById("txtNumber").value = "";
         document.getElementById("txtStreetName").value = "";
         document.getElementById("txtPostCode").value = "";


         var addressIndex = document.getElementById("selectLookup").selectedIndex;
         var addressSelection = document.getElementById("selectLookup").options[addressIndex].value

         var addressIndex_text = document.getElementById("selectLookup").selectedIndex;
         var addressSelection_text = document.getElementById("selectLookup").options[addressIndex_text].text

         //assign the UPRN to hidden field
         document.getElementById("hiddenUPRN").value = addressSelection;
         //assign the Address fields - display back inthe black box above search result
         document.getElementById("hiddenAddress").value = addressSelection_text;

         addressSelection_text = addressSelection_text.split(",");

         addrElements = ["txtHouseNumber", "txtBuildingName", "txtNumber", "txtStreetName", "txtPostCode"];
         var iAddrElement = 0;
         var i2 = 0;

         //assign Address texts fields
         if (addressSelection_text[0] != "Select your Address") {

             try {
                 for (var i = 0; i <= addressSelection_text.length; i++) {


                     if (addressSelection_text[i] != null && addressSelection_text[i] !== ' ALL') {
                         if (iAddrElement <= 4) {
                             document.getElementById(addrElements[iAddrElement]).value = addressSelection_text[i];
                         }
                         iAddrElement += 1;
                         i2 = i;
                     }
                 }

                 document.getElementById("txtLookupPostCode").value = addressSelection_text[i2];
             }
             catch (err) { alert(err); }
         }



     }


     //Validate lookup postcode
     if (typeof window.onload == 'function') { formatPostCode = window.onload; }
     window.onload = function () {
         if (window.formatPostCode) formatPostCode();

         runFormat();


     }

     /*FORMAT AND VALIDATE POSTCODE  - insert space and convert to upprcase*/
     function formatPostCode_begin() {

         var pcode = document.getElementById('txtLookupPostCode').value;
         //var pcode_2 = document.getElementById('txtPostCode').value;



         var patt = /^[a-z]{1,2}[0-9][0-9a-z]? ?[0-9][a-z]{2}$/i;

         var value = patt.test(pcode);

         if (!value) return "";

         patt = / /;
         value = patt.test(pcode);
         if (!value) // no space
         {
             var pcl = pcode.length;
             var pc1 = pcode.substring(0, pcl - 3);
             var pc2 = pcode.substring(pcl - 3, pcl - 0);
             pcode = pc1 + " " + pc2;
             pcode = pcode.toUpperCase();

             document.getElementById('txtLookupPostCode').value = pcode;


         }

         return pcode;
     }
     function runFormat() {
         document.getElementById('txtLookupPostCode').onkeyup = formatPostCode_begin;

     }
     //end postcode validation



     //Validate Address postcode
     if (typeof window.onload == 'function') { formatPostCode_2 = window.onload; }
     window.onload = function () {
         if (window.formatPostCode_2) formatPostCode_2();

         runFormat_2();
         setListLable();

         document.getElementById('txtPostCode').editable = false;

         document.getElementById("txtHouseNumber").disabled = true;
         document.getElementById("txtBuildingName").disabled = true;
         document.getElementById("txtNumber").disabled = true;
         document.getElementById("txtStreetName").disabled = true;
         document.getElementById("txtPostCode").disabled = true;
     }



     /*FORMAT AND VALIDATE POSTCODE  - insert space and convert to upprcase*/
     function formatPostCode_begin_2() {

         var pcode = document.getElementById('txtPostCode').value;
         //var pcode_2 = document.getElementById('txtPostCode').value;



         var patt = /^[a-z]{1,2}[0-9][0-9a-z]? ?[0-9][a-z]{2}$/i;

         var value = patt.test(pcode);

         if (!value) return "";

         patt = / /;
         value = patt.test(pcode);
         if (!value) // no space
         {
             var pcl = pcode.length;
             var pc1 = pcode.substring(0, pcl - 3);
             var pc2 = pcode.substring(pcl - 3, pcl - 0);
             pcode = pc1 + " " + pc2;
             pcode = pcode.toUpperCase();

             document.getElementById('txtPostCode').value = pcode;
         }

         return pcode;
     }
     function runFormat_2() {
         document.getElementById('txtPostCode').onkeyup = formatPostCode_begin_2;

     }


     function setListLable() {

         var lookup_list = document.getElementById('selectLookup');

         if (lookup_list.options.length > 1) {
             document.getElementById('listLabel').innerHTML = "&laquo Please select your Address &raquo; &nbsp; ";
         }
         if (lookup_list.options.length <= 1) {
             document.getElementById('listLabel').innerHTM = "";
         }
     }


     //make sure the postcode is entered
     function blockEmptyPostcodeSubmissions() {

         if (document.getElementById('txtLookupPostCode').value == "") {

             document.getElementById('Error_Field').innerHTML = "Please enter a valid postcode!";
             document.getElementById('txtLookupPostCode').focus();

             return false;
         }
         else {
             document.getElementById('Error_Field').innerHTML = "";
             return true;
         }

     }

     function blockEmptyDateRequest() {
         if (document.getElementById('hiddenUPRN').value == "" || document.getElementById('hiddenUPRN').value == "no") {

             document.getElementById('Error_Field').innerHTML = "Please provide a valid Address";
             document.getElementById('txtLookupPostCode').focus();

             return false;
         }
         else {
             document.getElementById('Error_Field').innerHTML = "";
             return true;
         }

     }

     function calendar_Button_Click() {

         var sCC;

         sCC = document.getElementById("hiddenCollectionCode").value;

         switch (sCC) {

             case "M1":
             case "M2":
             case "TU1":
             case "TU2":
             case "W1":
             case "W2":
             case "TH1":
             case "TH2":
             case "F1":
             case "F2":

                 window.location = document.getElementById("hiddenDocPath").value +
                                        document.getElementById("hiddenDocName").value +
                                        document.getElementById("hiddenCollectionCode").value +
                                        document.getElementById("hiddenDocExt").value;

                 break;

                 return true;

             default:

                 alert("Unfortunately, no calendar is available for the selected Address.\n\nPlease refer to the schedule below.");

                 return false;

         }

     }

</script>


 <div id="Error_Field" style="color: Red;">
 </div>

   

     

<div>

     				<div id="headingArea">

                   <p>Most households have recycling and rubbish collected on alternate weeks: recycling and food waste one week; rubbish and food waste the next week. Check the table below to find out when your collections will take place.<br><br>Please note this table does not take into account Christmas holidays.</p>

				</div>


<!-- ############### TEMPORARY WORDING FOR XMAS PERIOD ONLY
<div><h2>
Collection dates over Christmas and New Year have changed<br><a href="https://www.ealing.gov.uk/info/200950/recycling_and_christmas/1021/changes_to_refuse_and_recycling_collections">View revised dates</a>
</h2></div>
#################### END OF TEMPORARY WORDING FOR XMAS PERIOD ONLY
-->


<form method="post" action="/site/custom_scripts/waste_collection/waste_collection.aspx" id="Form1">
<div class="aspNetHidden">
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwULLTE0MDA5ODIxMjcPZBYEAgYPPCsACwBkAggPPCsACwBkZK2KWacxNMWb18vOTGFo9oxYaEiwhyI4d/et0yu+s57/" />
</div>

<div class="aspNetHidden">

	<input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="6FEC5A59" />
	<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEdAAMul8eokAuJuzesloktzbGxfM0kVSc4TD24D/S+ykhZviO9DoSWMLRMAQqZ6R0Abnb6jN0i9Pt4Mp8wKORLoZH/mbUjSGogX7q/zbJOh338tQ==" />
</div>

<table border="0">

   <tr><td colspan="2">
		<b>
			Enter your postcode and select your Address from the dropdown menu.<br>
			If there are multiple properties at your Address please ensure you select your specific flat,<br>
			you may have to scroll to the bottom of the list.
		</b>
	</td></tr>

   <tr><td colspan="2" align="left">
   <input type="hidden" name="hiddenLookup" value="LOOKUP">

  <div class="input">
   <input name="txtLookupPostCode" type="text" id="txtLookupPostCode" style="margin-left:0px" value="W13 9YU" />
   &nbsp;<input type="submit" class="button button--primary" value="Address lookup" id="btnAddressLookup" name="AddressLookup" onclick="return blockEmptyPostcodeSubmissions();" />
  </div>
  
  
   </td></tr>

   <tr><td style="padding-bottom:10px" colspan="2" align="left">

   <div id="listLabel" style="font-weight: bold; font-size:1em; color: red; padding-left:20px"></div>

  <div class="input">
   <select name="selectLookup" id="selectLookup" onchange="getSelectedAddress()">
	<option value="12125514        ">DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002950        ">1, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002961        ">2, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002972        ">3, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002982        ">4, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002993        ">5, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003002        ">6, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003003        ">7, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003004        ">8, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003005        ">9, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002951        ">10, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002952        ">11, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002953        ">12, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002954        ">13, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002955        ">14, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002956        ">15, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002957        ">16, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002958        ">17, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002959        ">18, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002960        ">19, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002962        ">20, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002963        ">21, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002964        ">22, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002965        ">23, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002966        ">24, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002967        ">25, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002968        ">26, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002969        ">27, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002970        ">28, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002971        ">29, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002973        ">30, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002974        ">31, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002975        ">32, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003159        ">33, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002976        ">34, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002977        ">35, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002978        ">36, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002979        ">37, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002980        ">38, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002981        ">39, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002983        ">40, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002984        ">41, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002985        ">42, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002986        ">43, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002987        ">44, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002988        ">45, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002989        ">46, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002990        ">47, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002991        ">48, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002992        ">49, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002994        ">50, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002995        ">51, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002996        ">52, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002997        ">53, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002998        ">54, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12002999        ">55, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003000        ">56, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
	<option value="12003001        ">57, DEAN COURT, BOWMANS CLOSE, ALL, WEST EALING, W13 9YU</option>
</select>
  </div>



   </td></tr>


 </table>
</form>
 

 <div style="height:20px;">
 </div>

<form name="waste_collection_getData" method="post" action="/site/custom_scripts/waste_collection/waste_collection.aspx">

 <table border="0">



   <tr><td colspan="2" align="center">

   <input type="hidden" id="hiddenLookup" name ="hiddenLookup" value="COLLECTION" />
   <input type="hidden" name ="hiddenUPRN" id="hiddenUPRN"  />
   <input type="hidden" name ="hiddenAddress" id="hiddenAddress"  />
   <input type="hidden" name ="hiddenTempScheduleAvailable" id="hiddenTempScheduleAvailable"  />

   <input name="hiddenCollectionCode" type="hidden" id="hiddenCollectionCode" />
   <input name="hiddenDocPath" type="hidden" id="hiddenDocPath" value="https://www.ealing.gov.uk/site/custom_scripts/waste_collection/docs/" />
   <input name="hiddenDocName" type="hidden" id="hiddenDocName" value="WC_Calendar_1218_" />
   <input name="hiddenDocExt" type="hidden" id="hiddenDocExt" value=".pdf" />
   <input name="hiddenTimedCollectionCode" type="hidden" id="hiddenTimedCollectionCode" />

   <input type="text" name="txtHouseNumber" id="txtHouseNumber" style="width:95%; text-align: center; font-size: 1em; font-weight: bold; color:#000000; background:#CCFFFF;" /></td></tr>

   <tr><td colspan="2" align="center">
   <input type="text" name="txtBuildingName" id="txtBuildingName" style="width:95%; text-align: center; font-size: 1em;  font-weight: bold; color:#000000; background:#CCFFFF;" /></td></tr>

   <tr><td colspan="2" align="center">
   <input type="text" name="txtNumber" id="txtNumber"  style="width:95%; text-align: center; font-size: 1em;  font-weight: bold; color:#000000; background:#CCFFFF;" /></td></tr>

   <tr><td colspan="2" align="center"><input type="text" name="txtStreetName"  id="txtStreetName" style="width:95%; text-align: center; font-size: 1em;  font-weight: bold; color:#000000; background:#CCFFFF;" /></td></tr>
   <tr><td colspan="2" align="center"><input type="text" name="txtPostCode" id="txtPostCode" style="width:95%; text-align: center; font-size: 1em;  font-weight: bold; color:#000000; background:#CCFFFF;" />
   </td></tr>



   <tr><td style="padding-top:20px; padding-bottom:10px" colspan="2" align="center"><input  type="submit" class="button button--primary" onclick="return blockEmptyDateRequest();" value="Show my collection day" /></td></tr>

   </table>
  </form>
  


    

    


</div>

<!-- ################ END OF WASTE COLLECTION FORM ############ -->

<!-- ###################################### -->
<div class="supplements supplements--secondary">
</div>
        </div>

                <div class="sidebar sidebar--primary">


                <div class="supplements supplements--tertiary">

</div>

            </div>


        </div>
    </main>
        <!-- googleoff: index -->
        <footer class="site-footer" role="contentinfo">
            <div class="container">
                <div class="site-footer__social-links">
                    <h2>Follow us</h2>
                    <ul class="item-list item-list--inline">
                        <li><a class="youtube" href="https://www.youtube.com/channel/UCoI9fzIsjNbPOOMnTO6fxhQ">YouTube</a></li>
                        <li><a class="twitter" href="https://twitter.com/ealingcouncil">Twitter</a></li>
                        <li><a class="flickr" href="https://www.flickr.com/photos/ealingaltogether/">Flickr</a></li>
                        <li><a class="rss" href="https://www.ealing.gov.uk/rss/news">RSS</a></li>
                    </ul>
                </div>
                <ul class="site-footer__links item-list item-list--inline">
                    <li>
                        <a accesskey="1" href="/site/scripts/terms.php">Terms and disclaimer</a>
                    </li>
                    <li>
                        <a accesskey="2" href="/site/scripts/site_map.php" rel="nofollow">Site map</a>
                    </li>
                    <li>
                        <a accesskey="3" href="/accessibility/settings">Accessibility</a>
                    </li>
                    <li>
                        <a accesskey="4" href="/info/201045/data_protection/1420/privacy_statement" rel="nofollow">Privacy</a>
                    </li>
                    <li>
                        <a accesskey="4" href="/homepage/24/about_us" rel="nofollow">About us</a>
                    </li>                    
                    <!--<li>
                        <a href="/forms/form/86/en/report_problem#top" rel="nofollow">Back to the top</a>
                    </li>
                    <li>
                        <a href="/feedback">Feedback</a>
                    </li>
                    <li>
                        <a href="/statistics">Statistics</a>
                    </li>
                    <li>
                        <a href="/page_comments/L2Zvcm1zL2Zvcm0vODYvZW4vcmVwb3J0X3Byb2JsZW0=">Comment on this page</a>
                    </li>
                    <li>
                        <a rel="nofollow" href="#" onclick="window.print(); return false;">Print this page</a>
                    </li>-->
                </ul>

                </p>
                <!--<ul class="visuallyhidden">
                    <li>
                        <a accesskey="1" href="/" rel="nofollow">Homepage</a>
                    </li>
                    <li>
                        <a accesskey="2" href="/whats_new" rel="nofollow">What's new</a>
                    </li>
                    <li>
                        <a accesskey="4" href="/site_search" rel="nofollow">Search facility</a>
                    </li>
                    <li>
                        <a accesskey="5" href="/faqs" rel="nofollow">Frequently asked questions</a>
                    </li>
                    <li>
                        <a accesskey="6" href="/a_to_z" rel="nofollow">Help</a>
                    </li>
                    <li>
                        <a accesskey="7" href="/contact" rel="nofollow">Contact details</a>
                    </li>
                    <li>
                        <a accesskey="9" href="/feedback" rel="nofollow">Feedback</a>
                    </li>
                    <li>
                        <a accesskey="0" href="/accessibility" rel="nofollow">Access key details</a>
                    </li>
                </ul>-->
            </div>
        </footer>
        <!-- googleon: index -->
        <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
        <script>window.jQuery || document.write('<script src="/site/libs/jquery/dist/jquery.min.js"><\/script>')</script>
        <script src="/site/dist/ealing.js"></script>
        <!-- SiteImprove analytics -->
        <script type="text/javascript">/*<![CDATA[*/(function() {var sz = document.createElement('script'); sz.type = 'text/javascript'; sz.async = true;sz.src = '//uk1.siteimprove.com/js/siteanalyze_432813.js';var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(sz, s);})();/*]]>*/</script>
    </body>
</html>`