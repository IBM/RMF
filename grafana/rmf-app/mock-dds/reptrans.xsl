<?xml version="1.0"?>
<!--
  ~ Licensed materials - Property of IBM
  ~ 5655-ZOS
  ~ (c) Copyright IBM Corp. 2023 All Rights Reserved.
  ~ (c) Copyright Rocket Software, Inc. 2023 All Rights Reserved.
  ~ US Government Users Restricted Rights - Use, duplication or
  ~ disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
  -->
<!--****************************************************************-->
<!-- Name: reptrans.xsl                                             -->
<!--                                                                -->
<!-- Function: Translate columns headers to readable text           -->
<!--                                                                -->
<!--****************************************************************-->
<xsl:stylesheet
        xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
        version="1.0">
    <!-- ************************************************************ -->
    <!-- Create buttons header                                        -->
    <!-- ************************************************************ -->
    <xsl:template name="hdrtrans">
        <xsl:param name="report" select="'NA'"/>
        <xsl:param name="var" select="."/>
        <xsl:choose>
            <!-- SYSRG -->
            <xsl:when test="$var = 'SRGNAME'">Resource Group Name</xsl:when>
            <xsl:when test="$var = 'SRGTYPE'">Type</xsl:when>
            <xsl:when test="$var = 'SRGSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'SRGSCTRC'">Associated WLM Class</xsl:when>
            <xsl:when test="$var = 'SRGSRTYP'">Capacity Def</xsl:when>
            <xsl:when test="$var = 'SRGSRMIN'">Capacity Min</xsl:when>
            <xsl:when test="$var = 'SRGSRMAX'">Capacity Max</xsl:when>
            <xsl:when test="$var = 'SRGSRACP'">Capacity Actual (# CPs)</xsl:when>
            <xsl:when test="$var = 'SRGSRAMS'">Capacity Actual (MSU)</xsl:when>
            <xsl:when test="$var = 'SRGSRASU'">Capacity Actual (SU/sec)</xsl:when>
            <xsl:when test="$var = 'SRGSPINC'">Include Specialty Processors</xsl:when>
            <xsl:when test="$var = 'SRGMMAX'">Memory Limit</xsl:when>
            <xsl:when test="$var = 'SRGMACT'">Memory Actual</xsl:when>
            <xsl:when test="$var = 'SRGSID'">SMF ID</xsl:when>
            <xsl:when test="$var = 'SRGDESCD'">Description</xsl:when>
            <!-- CRYOVW/CRYPTO -->
            <xsl:when test="$var = 'CRYCTYPE'">Crypto Card Type</xsl:when>
            <xsl:when test="$var = 'CRYCIDX'">Crypto Card Index</xsl:when>
            <xsl:when test="$var = 'CRYCPC'">CPC Name</xsl:when>
            <xsl:when test="$var = 'CRYSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'CRYUDID'">Usage Domain ID</xsl:when>
            <xsl:when test="$var = 'CRYSCPM'">Scope</xsl:when>
            <xsl:when test="$var = 'CRYCMODE'">Cryptographic Mode</xsl:when>
            <xsl:when test="$var = 'CRYOPRT'">Total Rate</xsl:when>
            <xsl:when test="$var = 'CRYOPET'">Total Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRYTUTL'">Total Util %</xsl:when>
            <xsl:when test="$var = 'CRCKGORT'">RSA-Key-Gen Rate</xsl:when>
            <xsl:when test="$var = 'CRCKGOET'">RSA-Key-Gen Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRCKGUTL'">RSA-Key-Gen Util %</xsl:when>
            <xsl:when test="$var = 'CRARKLEN'">RSA Key Length</xsl:when>
            <xsl:when test="$var = 'CRAMOPRT'">ME-Format RSA Rate</xsl:when>
            <xsl:when test="$var = 'CRAMOPET'">ME-Format RSA Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRAMUTL'">ME-Format RSA Util %</xsl:when>
            <xsl:when test="$var = 'CRACOPRT'">CRT-Format RSA Rate</xsl:when>
            <xsl:when test="$var = 'CRACOPET'">CRT-Format RSA Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRACUTL'">CRT-Format RSA Util %</xsl:when>
            <xsl:when test="$var = 'CRPSART'">Slow Asym-Key Rate</xsl:when>
            <xsl:when test="$var = 'CRPSAET'">Slow Asym-Key Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRPSAUTL'">Slow Asym-Key Util %</xsl:when>
            <xsl:when test="$var = 'CRPFART'">Fast Asym-Key Rate</xsl:when>
            <xsl:when test="$var = 'CRPFAET'">Fast Asym-Key Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRPFAUTL'">Fast Asym-Key Util %</xsl:when>
            <xsl:when test="$var = 'CRPSPRT'">Sym-Key Partial Rate</xsl:when>
            <xsl:when test="$var = 'CRPSPET'">Sym-Key Partial Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRPSPUTL'">Sym-Key Partial Util %</xsl:when>
            <xsl:when test="$var = 'CRPSCRT'">Sym-Key Complete Rate</xsl:when>
            <xsl:when test="$var = 'CRPSCET'">Sym-Key Complete Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRPSCUTL'">Sym-Key Complete Util %</xsl:when>
            <xsl:when test="$var = 'CRPAGRT'">Asym-Key Generation Rate</xsl:when>
            <xsl:when test="$var = 'CRPAGET'">Asym-Key Generation Avg Exec Time</xsl:when>
            <xsl:when test="$var = 'CRPAGUTL'">Asym-Key Generation Util %</xsl:when>
            <xsl:when test="$var = 'CRYSID'">SMF ID</xsl:when>
            <!-- EADM (aka SCM) -->
            <xsl:when test="$var = 'SCMHSCH'">Total Number of SSCH</xsl:when>
            <xsl:when test="$var = 'SCMHSCR'">SSCH Rate</xsl:when>
            <xsl:when test="$var = 'SCMHFPT'">Avg Function Pending Time</xsl:when>
            <xsl:when test="$var = 'SCMHIQT'">Avg IOP Queue Time</xsl:when>
            <xsl:when test="$var = 'SCMHCRT'">Avg Initial Cmd Response Time</xsl:when>
            <xsl:when test="$var = 'SCMHRRC'">Compression Request Rate</xsl:when>
            <xsl:when test="$var = 'SCMHTPC'">Compression Throughput</xsl:when>
            <xsl:when test="$var = 'SCMHRCC'">Compression Ratio</xsl:when>
            <xsl:when test="$var = 'SCMHRRD'">Decompression Request Rate</xsl:when>
            <xsl:when test="$var = 'SCMHTPD'">Decompression Throughput</xsl:when>
            <xsl:when test="$var = 'SCMHRCD'">Decompression Ratio</xsl:when>
            <xsl:when test="$var = 'SCMRPID'">Card ID</xsl:when>
            <xsl:when test="$var = 'SCMUTL'">Util% (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMUTLC'">Util% (Total)</xsl:when>
            <xsl:when test="$var = 'SCMDRD'">Read B/Sec (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMDRDC'">Read B/Sec (Total)</xsl:when>
            <xsl:when test="$var = 'SCMDWR'">Write B/Sec (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMDWRC'">Write B/Sec (Total)</xsl:when>
            <xsl:when test="$var = 'SCMRQR'">Request Rate (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMRQRC'">Request Rate (Total)</xsl:when>
            <xsl:when test="$var = 'SCMART'">Avg Response Time (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMARTC'">Avg Response Time (Total)</xsl:when>
            <xsl:when test="$var = 'SCMAQTC'">Avg IOP Queue Time (Total)</xsl:when>
            <xsl:when test="$var = 'SCMTRQ'">Requests (LPAR)</xsl:when>
            <xsl:when test="$var = 'SCMTRQC'">Requests (Total)</xsl:when>
            <!-- ZFSOVW -->
            <xsl:when test="$var = 'ZFOPSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'ZFOPIORP'">Avg Response Time I/O%</xsl:when>
            <xsl:when test="$var = 'ZFOPLORP'">Avg Response Time Lock%</xsl:when>
            <xsl:when test="$var = 'ZFOPSLRP'">Avg Response Time Sleep%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRT'">User Cache Request Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPUCHP'">User Cache Hit%</xsl:when>
            <xsl:when test="$var = 'ZFOPVCRT'">Vnode Cache Request Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPVCHP'">Vnode Cache Hit%</xsl:when>
            <xsl:when test="$var = 'ZFOPMCRT'">Metadata Cache Request Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPMCHP'">Metadata Cache Hit%</xsl:when>
            <!-- I/O cache details -->
            <xsl:when test="$var = 'ZFOPITY1'">I/O Details Type (1)</xsl:when>
            <xsl:when test="$var = 'ZFOPICT1'">I/O Details #&#160;Requests (1)</xsl:when>
            <xsl:when test="$var = 'ZFOPIWT1'">I/O Details #&#160;Waits (1)</xsl:when>
            <xsl:when test="$var = 'ZFOPICA1'">I/O Details #&#160;Cancels (1)</xsl:when>
            <xsl:when test="$var = 'ZFOPIMG1'">I/O Details #&#160;Merges (1)</xsl:when>
            <xsl:when test="$var = 'ZFOPITY2'">I/O Details Type (2)</xsl:when>
            <xsl:when test="$var = 'ZFOPICT2'">I/O Details #&#160;Requests (2)</xsl:when>
            <xsl:when test="$var = 'ZFOPIWT2'">I/O Details #&#160;Waits (2)</xsl:when>
            <xsl:when test="$var = 'ZFOPICA2'">I/O Details #&#160;Cancels (2)</xsl:when>
            <xsl:when test="$var = 'ZFOPIMG2'">I/O Details #&#160;Merges (2)</xsl:when>
            <xsl:when test="$var = 'ZFOPITY3'">I/O Details Type (3)</xsl:when>
            <xsl:when test="$var = 'ZFOPICT3'">I/O Details #&#160;Requests (3)</xsl:when>
            <xsl:when test="$var = 'ZFOPIWT3'">I/O Details #&#160;Waits (3)</xsl:when>
            <xsl:when test="$var = 'ZFOPICA3'">I/O Details #&#160;Cancels (3)</xsl:when>
            <xsl:when test="$var = 'ZFOPIMG3'">I/O Details #&#160;Merges (3)</xsl:when>
            <!-- User cache details -->
            <xsl:when test="$var = 'ZFOPUCSZ'">User Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFOPUCSF'">User Cache Storage fixed</xsl:when>
            <xsl:when test="$var = 'ZFOPUCTP'">User Cache Total Pages</xsl:when>
            <xsl:when test="$var = 'ZFOPUCFP'">User Cache Free Pages</xsl:when>
            <xsl:when test="$var = 'ZFOPUCSG'">User Cache Segments</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRR'">User Cache Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRH'">User Cache Read Hit%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRD'">User Cache Read Dly%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCAR'">User Cache Async Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPUCWR'">User Cache Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPUCWH'">User Cache Write Hit%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCWD'">User Cache Write Dly%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCSW'">User Cache Scheduled Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRP'">User Cache Read%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCDP'">User Cache Dly%</xsl:when>
            <xsl:when test="$var = 'ZFOPUCRW'">User Cache Page Reclaim Writes</xsl:when>
            <xsl:when test="$var = 'ZFOPUCFS'">User Cache Fsynchs</xsl:when>
            <!-- Vnode cache details -->
            <xsl:when test="$var = 'ZFOPVCSZ'">Vnode Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFOPVCAL'"># Allocated Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFOPVCSN'">Vnode Structure Size</xsl:when>
            <xsl:when test="$var = 'ZFOPVCEX'"># Extended Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFOPVCSE'">Extended Vnode Structure Size</xsl:when>
            <xsl:when test="$var = 'ZFOPVCOP'"># Open Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFOPVCHE'"># USS Held Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFOPVCRQ'">Vnode Cache #&#160;Requests</xsl:when>
            <xsl:when test="$var = 'ZFOPVCCR'">Vnode Cache #&#160;Allocs</xsl:when>
            <xsl:when test="$var = 'ZFOPVCDL'">Vnode Cache #&#160;Deletes</xsl:when>
            <!-- Metadata cache details -->
            <xsl:when test="$var = 'ZFOPMCSZ'">Metadata Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFOPMCSF'">Metadata Cache Storage fixed</xsl:when>
            <xsl:when test="$var = 'ZFOPMCBU'">Metadata Cache #&#160;Buffers</xsl:when>
            <xsl:when test="$var = 'ZFOPMCRQ'">Metadata Cache #&#160;Requests</xsl:when>
            <xsl:when test="$var = 'ZFOPMCUD'">Metadata Cache #&#160;Updates</xsl:when>
            <xsl:when test="$var = 'ZFOPMCPW'">Metadata Cache #&#160;Partial Writes</xsl:when>
            <!-- ZFSFS -->
            <xsl:when test="$var = 'ZFFPFSN'">File System Name</xsl:when>
            <xsl:when test="$var = 'ZFFPSYSC'">Connected System</xsl:when>
            <xsl:when test="$var = 'ZFFPSYSO'">Owner</xsl:when>
            <xsl:when test="$var = 'ZFFPMODE'">File System Mode</xsl:when>
            <xsl:when test="$var = 'ZFFPSIZE'">Size</xsl:when>
            <xsl:when test="$var = 'ZFFPUSGP'">Usg%</xsl:when>
            <xsl:when test="$var = 'ZFFPAPIR'">I/O Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPAPRT'">Response Time</xsl:when>
            <xsl:when test="$var = 'ZFFPAPRP'">Read%</xsl:when>
            <xsl:when test="$var = 'ZFFPAPXR'">XCF Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPFSMP'">Mount Point</xsl:when>
            <xsl:when test="$var = 'ZFFPFSVN'"># Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFFPFSVU'"># USS Held Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFFPOBJO'"># Open Objects</xsl:when>
            <xsl:when test="$var = 'ZFFPOBJT'"># Tokens</xsl:when>
            <xsl:when test="$var = 'ZFFPFSUC'"># User Cache 4K Pages</xsl:when>
            <xsl:when test="$var = 'ZFFPFSMC'"># Metadata Cache 8K Pages</xsl:when>
            <xsl:when test="$var = 'ZFFPAPRR'">Application Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPARRT'">Application Read Response Time</xsl:when>
            <xsl:when test="$var = 'ZFFPXFRR'">XCF Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPXRRT'">XCF Read Response Time</xsl:when>
            <xsl:when test="$var = 'ZFFPIORR'">Aggregate Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPAPWR'">Application Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPAWRT'">Application Write Response Time</xsl:when>
            <xsl:when test="$var = 'ZFFPXFWR'">XCF Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPXWRT'">XCF Write Response Time</xsl:when>
            <xsl:when test="$var = 'ZFFPIOWR'">Aggregate Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFFPESPC'"># ENOSPC Errors</xsl:when>
            <xsl:when test="$var = 'ZFFPEDIO'"># Disk I/O Errors</xsl:when>
            <xsl:when test="$var = 'ZFFPEXCF'"># XCF Communication Failures</xsl:when>
            <xsl:when test="$var = 'ZFFPOPCA'"># Cancelled Operations</xsl:when>
            <!-- ZFSKN -->
            <xsl:when test="$var = 'ZFKPSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'ZFKPRQRL'">Request Rate Local</xsl:when>
            <xsl:when test="$var = 'ZFKPRQRR'">Request Rate Remote</xsl:when>
            <xsl:when test="$var = 'ZFKPXFRL'">XCF Rate Local</xsl:when>
            <xsl:when test="$var = 'ZFKPXFRR'">XCF Rate Remote</xsl:when>
            <xsl:when test="$var = 'ZFKPRPTL'">Response Time Local</xsl:when>
            <xsl:when test="$var = 'ZFKPRPTR'">Response Time Remote</xsl:when>
            <!-- I/O cache details -->
            <xsl:when test="$var = 'ZFSICNT1'">Number of Requests (Type 1)</xsl:when>
            <xsl:when test="$var = 'ZFSIWT01'">Requests waiting for I/O Completion (Type 1)</xsl:when>
            <xsl:when test="$var = 'ZFSICAN1'">Requests cancelled (Type 1)</xsl:when>
            <xsl:when test="$var = 'ZFSIMRG1'">Requests merged (Type 1)</xsl:when>
            <xsl:when test="$var = 'ZFSITYP1'">Type Description (Type 1)</xsl:when>
            <xsl:when test="$var = 'ZFSICNT2'">Number of Requests (Type 2)</xsl:when>
            <xsl:when test="$var = 'ZFSIWT02'">Requests waiting for I/O Completion (Type 2)</xsl:when>
            <xsl:when test="$var = 'ZFSICAN2'">Requests cancelled (Type 2)</xsl:when>
            <xsl:when test="$var = 'ZFSIMRG2'">Requests merged (Type 2)</xsl:when>
            <xsl:when test="$var = 'ZFSITYP2'">Type Description (Type 2)</xsl:when>
            <xsl:when test="$var = 'ZFSICNT3'">Number of Requests (Type 3)</xsl:when>
            <xsl:when test="$var = 'ZFSIWT03'">Requests waiting for I/O Completion (Type 3)</xsl:when>
            <xsl:when test="$var = 'ZFSICAN3'">Requests cancelled (Type 3)</xsl:when>
            <xsl:when test="$var = 'ZFSIMRG3'">Requests merged (Type 3)</xsl:when>
            <xsl:when test="$var = 'ZFSITYP3'">Type Description (Type 3)</xsl:when>
            <!-- User cache details -->
            <xsl:when test="$var = 'ZFSURDRT'">Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFSUWRRT'">Write Rate</xsl:when>
            <xsl:when test="$var = 'ZFSURDHT'">Read Requests completed without I/O</xsl:when>
            <xsl:when test="$var = 'ZFSUWRHT'">Write Requests completed without I/O</xsl:when>
            <xsl:when test="$var = 'ZFSURDDY'">Percentage of delayed Reads</xsl:when>
            <xsl:when test="$var = 'ZFSUWRDY'">Percentage of delayed Writes</xsl:when>
            <xsl:when test="$var = 'ZFSUARRT'">Number of Read-aheads</xsl:when>
            <xsl:when test="$var = 'ZFSUSWRT'">Number of scheduled Writes</xsl:when>
            <xsl:when test="$var = 'ZFSUPRWT'">Number of page reclaim Writes</xsl:when>
            <xsl:when test="$var = 'ZFSUFSYN'">Number of fsynchs</xsl:when>
            <xsl:when test="$var = 'ZFSUSIZE'">Total Size of User Cache</xsl:when>
            <xsl:when test="$var = 'ZFSUTOPG'">Number of Pages in User Cache</xsl:when>
            <xsl:when test="$var = 'ZFSUFRPG'">Number of free Pages in User Cache</xsl:when>
            <xsl:when test="$var = 'ZFSUSEGM'">Number of allocated Segments</xsl:when>
            <xsl:when test="$var = 'ZFSURDAH'">Read-ahead</xsl:when>
            <xsl:when test="$var = 'ZFSUSTOF'">User Cache Storage fixed</xsl:when>
            <!-- Vnode cache details -->
            <xsl:when test="$var = 'ZFSVRQRT'">Request Rate to Vnode Cache</xsl:when>
            <xsl:when test="$var = 'ZFSVHIT'">Request Hits in Vnode Cache</xsl:when>
            <xsl:when test="$var = 'ZFSVNODE'">Vnode Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFSVSSIZ'">Vnode Structure Size</xsl:when>
            <xsl:when test="$var = 'ZFSVNEXT'">Extended Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFSVESIZ'">Extended Vnode Size</xsl:when>
            <xsl:when test="$var = 'ZFSVNOPN'">Open Vnodes</xsl:when>
            <xsl:when test="$var = 'ZFSVNHLD'">Held Vnodes by USS</xsl:when>
            <!-- Metadata cache details -->
            <xsl:when test="$var = 'ZFSMRQRT'">Request Rate to Metadata Cache</xsl:when>
            <xsl:when test="$var = 'ZFSMHIT'">Requests to Metadata Cache completed without I/O</xsl:when>
            <xsl:when test="$var = 'ZFSMSIZE'">Metadata Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFSMBUFF'">Number of Buffers (8k) in Metadata Cache</xsl:when>
            <xsl:when test="$var = 'ZFSMSTOF'">Metadata Cache Storage fixed</xsl:when>
            <xsl:when test="$var = 'ZFSBRQRT'">Request Rate to Metadata Backing Cache</xsl:when>
            <xsl:when test="$var = 'ZFSBHIT'">Requests to Metadata Backing Cache completed without I/O</xsl:when>
            <xsl:when test="$var = 'ZFSBSIZE'">Metadata Backing Cache Size</xsl:when>
            <xsl:when test="$var = 'ZFSBBUFF'">Number of Buffers (8k) in Metadata Backing Cache</xsl:when>
            <xsl:when test="$var = 'ZFSBSTOF'">Metadata Backing Cache Storage fixed</xsl:when>
            <!-- Transaction cache details -->
            <xsl:when test="$var = 'ZFSTRQRT'">Transactions started in the Transaction Cache</xsl:when>
            <xsl:when test="$var = 'ZFSTMERG'">Merged Transactions to Transaction Class</xsl:when>
            <xsl:when test="$var = 'ZFSTACTV'">Active Transactions</xsl:when>
            <xsl:when test="$var = 'ZFSTPEND'">Pending Transactions</xsl:when>
            <xsl:when test="$var = 'ZFSTCOMP'">Complete Transactions</xsl:when>
            <xsl:when test="$var = 'ZFSTFREE'">Free Transactions</xsl:when>
            <xsl:when test="$var = 'ZFSTALLO'">Transaction Structures allocated in the Transaction Cache</xsl:when>
            <xsl:when test="$var = 'ZFSPAGG'">Aggregate Name</xsl:when>
            <xsl:when test="$var = 'ZFSPASIZ'">Aggregate Size</xsl:when>
            <xsl:when test="$var = 'ZFSPAUSE'">% Aggregate Use</xsl:when>
            <xsl:when test="$var = 'ZFSPAMD1'">Aggregate Mode (R/W, R/O)</xsl:when>
            <xsl:when test="$var = 'ZFSPAMD2'">Aggregate Mode (MS, CP)</xsl:when>
            <xsl:when test="$var = 'ZFSPAFS'">Number of File Systems</xsl:when>
            <xsl:when test="$var = 'ZFSPARR'">Aggregate Read Rate</xsl:when>
            <xsl:when test="$var = 'ZFSPAWR'">Aggregate Write Rate</xsl:when>
            <!-- SPACEG -->
            <xsl:when test="$var = 'SPGPSGN'">Storage Group Name</xsl:when>
            <xsl:when test="$var = 'SPGPTSP'">Total Capacity (MB)</xsl:when>
            <xsl:when test="$var = 'SPGPFSP'">Free Space (MB)</xsl:when>
            <xsl:when test="$var = 'SPGPFSR'">Free Space %</xsl:when>
            <xsl:when test="$var = 'SPGPNVO'">Number of Volumes</xsl:when>
            <xsl:when test="$var = 'SPGPUVO'">Unallocated Volumes</xsl:when>
            <!-- SPACED -->
            <xsl:when test="$var = 'SPDPVOL'">Volume</xsl:when>
            <xsl:when test="$var = 'SPDPTSP'">Total Capacity (MB)</xsl:when>
            <xsl:when test="$var = 'SPDPFSP'">Free Space (MB)</xsl:when>
            <xsl:when test="$var = 'SPDPFSR'">Free Space %</xsl:when>
            <xsl:when test="$var = 'SPDPLBK'">Largest Block (MB)</xsl:when>
            <xsl:when test="$var = 'SPDPSGN'">Storage Group Name</xsl:when>
            <!-- OPD -->
            <xsl:when test="$var = 'OPDHPROC'">Kernel Procedure</xsl:when>
            <xsl:when test="$var = 'OPDHASID'">Kernel ASID</xsl:when>
            <xsl:when test="$var = 'OPDHOPTN'">Option</xsl:when>
            <xsl:when test="$var = 'OPDHOPT'">Select</xsl:when>
            <xsl:when test="$var = 'OPDHPLST'">BPXPRM</xsl:when>
            <xsl:when test="$var = 'OPDPJOBN'">Job Name</xsl:when>
            <xsl:when test="$var = 'OPDPUSEN'">User</xsl:when>
            <xsl:when test="$var = 'OPDPASID'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'OPDPASIX'">ASID (hex)</xsl:when>
            <xsl:when test="$var = 'OPDPPRID'">Process ID</xsl:when>
            <xsl:when test="$var = 'OPDPPPID'">Parent Process ID</xsl:when>
            <xsl:when test="$var = 'OPDPLATW'">Latch Wait</xsl:when>
            <xsl:when test="$var = 'OPDPSTAT'">Process State</xsl:when>
            <xsl:when test="$var = 'OPDPAPPL'">Appl&#160;%</xsl:when>
            <xsl:when test="$var = 'OPDPTOT'">Total CPU Time</xsl:when>
            <xsl:when test="$var = 'OPDPSERV'">Server Type</xsl:when>
            <xsl:when test="$var = 'OPDICMD'">Command</xsl:when>
            <xsl:when test="$var = 'OPDIACTF'">Active Files</xsl:when>
            <xsl:when test="$var = 'OPDIMAXF'">Max. Files</xsl:when>
            <xsl:when test="$var = 'OPDITIMD'">Start&#160;Time&#160;and&#160;Date</xsl:when>
            <xsl:when test="$var = 'OPDISTA1'">Process&#160;State&#160;Description&#160;1</xsl:when>
            <xsl:when test="$var = 'OPDISTA2'">Process&#160;State&#160;Description&#160;2</xsl:when>
            <xsl:when test="$var = 'OPDISTA3'">Process&#160;State&#160;Description&#160;3</xsl:when>
            <xsl:when test="$var = 'OPDISTA4'">Process&#160;State&#160;Description&#160;4</xsl:when>
            <xsl:when test="$var = 'OPDISTA5'">Process&#160;State&#160;Description&#160;5</xsl:when>
            <xsl:when test="$var = 'OPDPJID'">JES ID</xsl:when>
            <!-- ENCLAVE -->
            <xsl:when test="$var = 'ENCHSST'">Subsystem Type</xsl:when>
            <xsl:when test="$var = 'ENCHEON'">Enclave Owner</xsl:when>
            <xsl:when test="$var = 'ENCHCLS'">Class/Group</xsl:when>
            <xsl:when test="$var = 'ENCHAPP'">CP Appl %</xsl:when>
            <xsl:when test="$var = 'ENCHEAP'">EAppl %</xsl:when>
            <xsl:when test="$var = 'ENCHAPI'">AAP Appl %</xsl:when>
            <xsl:when test="$var = 'ENCHAPS'">IIP Appl %</xsl:when>
            <xsl:when test="$var = 'ENCENAME'">Enclave Name</xsl:when>
            <xsl:when test="$var = 'ENCCLASS'">CLS/GRP</xsl:when>
            <xsl:when test="$var = 'ENCGOAL'">Goal</xsl:when>
            <xsl:when test="$var = 'ENCGPERC'">Goal %</xsl:when>
            <xsl:when test="$var = 'ENCPER'">Period</xsl:when>
            <xsl:when test="$var = 'ENCDENC'">Dependent</xsl:when>
            <xsl:when test="$var = 'ENCXENC'">Multi System</xsl:when>
            <xsl:when test="$var = 'ENCTCPU'">Total CPU Time</xsl:when>
            <xsl:when test="$var = 'ENCTIFA'">Total AAP Time</xsl:when>
            <xsl:when test="$var = 'ENCTIFC'">Total AAP on CP Time</xsl:when>
            <xsl:when test="$var = 'ENCTSUP'">Total IIP Time</xsl:when>
            <xsl:when test="$var = 'ENCTSUC'">Total IIP on CP Time</xsl:when>
            <xsl:when test="$var = 'ENCDCPU'">Delta CPU Time</xsl:when>
            <xsl:when test="$var = 'ENCDIFA'">Delta AAP Time</xsl:when>
            <xsl:when test="$var = 'ENCDIFC'">Delta AAP on CP Time</xsl:when>
            <xsl:when test="$var = 'ENCDSUP'">Delta IIP Time</xsl:when>
            <xsl:when test="$var = 'ENCDSUC'">Delta IIP on CP Time</xsl:when>
            <xsl:when test="$var = 'ENCDCPUP'">Delta CPU %</xsl:when>
            <xsl:when test="$var = 'ENCDIFAP'">Delta AAP %</xsl:when>
            <xsl:when test="$var = 'ENCDIFCP'">Delta AAP on CP %</xsl:when>
            <xsl:when test="$var = 'ENCDSUPP'">Delta IIP %</xsl:when>
            <xsl:when test="$var = 'ENCDSUCP'">Delta IIP on CP %</xsl:when>
            <xsl:when test="$var = 'ENCSAMP'">Total Samples</xsl:when>
            <xsl:when test="$var = 'ENCTUSG'">Total Using %</xsl:when>
            <xsl:when test="$var = 'ENCTDLY'">Total Delay %</xsl:when>
            <xsl:when test="$var = 'ENCIDLE'">Total Idle %</xsl:when>
            <xsl:when test="$var = 'ENCCUSG'">CPU Using %</xsl:when>
            <xsl:when test="$var = 'ENCIFAU'">AAP Using %</xsl:when>
            <xsl:when test="$var = 'ENCIFCU'">AAP on CP Using %</xsl:when>
            <xsl:when test="$var = 'ENCSUPU'">IIP Using %</xsl:when>
            <xsl:when test="$var = 'ENCSUCU'">IIP on CP Using %</xsl:when>
            <xsl:when test="$var = 'ENCCDLY'">CPU Delay %</xsl:when>
            <xsl:when test="$var = 'ENCIFAD'">AAP Delay %</xsl:when>
            <xsl:when test="$var = 'ENCSUPD'">IIP Delay %</xsl:when>
            <xsl:when test="$var = 'ENCIUSG'">I/O Using %</xsl:when>
            <xsl:when test="$var = 'ENCIDLY'">I/O Delay %</xsl:when>
            <xsl:when test="$var = 'ENCCCAP'">CPU Capping %</xsl:when>
            <xsl:when test="$var = 'ENCSTOR'">Storage Delay %</xsl:when>
            <xsl:when test="$var = 'ENCUNKN'">Unknown Delay %</xsl:when>
            <xsl:when test="$var = 'ENCQUED'">Queue Delay %</xsl:when>
            <xsl:when test="$var = 'ENCESTYP'">Subsystem Type</xsl:when>
            <xsl:when test="$var = 'ENCEOWNM'">Owner</xsl:when>
            <xsl:when test="$var = 'ENCEOSYS'">Owner System</xsl:when>
            <xsl:when test="$var = 'ENCETOKN'">Enclave&#160;Token</xsl:when>
            <xsl:when test="$var = 'ENCXTOKN'">Export&#160;Token</xsl:when>
            <xsl:when test="$var = 'ENCATTN'">Number of Attributes</xsl:when>
            <xsl:when test="$var = 'ENCATT00'">Attribute</xsl:when>
            <xsl:when test="$var = 'ENCATT01'">Att: Account</xsl:when>
            <xsl:when test="$var = 'ENCATT02'">Att: Collection</xsl:when>
            <xsl:when test="$var = 'ENCATT03'">Att: CNCTN</xsl:when>
            <xsl:when test="$var = 'ENCATT04'">Att: Correlation</xsl:when>
            <xsl:when test="$var = 'ENCATT05'">Att: LU</xsl:when>
            <xsl:when test="$var = 'ENCATT06'">Att: NetID</xsl:when>
            <xsl:when test="$var = 'ENCATT07'">Att: Package</xsl:when>
            <xsl:when test="$var = 'ENCATT08'">Att: Planname</xsl:when>
            <xsl:when test="$var = 'ENCATT09'">Att: Procedure</xsl:when>
            <xsl:when test="$var = 'ENCATT10'">Att: Process</xsl:when>
            <xsl:when test="$var = 'ENCATT11'">Att: Sched Env</xsl:when>
            <xsl:when test="$var = 'ENCATT12'">Att: Subsys Collection</xsl:when>
            <xsl:when test="$var = 'ENCATT13'">Att: Subsys Name</xsl:when>
            <xsl:when test="$var = 'ENCATT14'">Att: SSPM</xsl:when>
            <xsl:when test="$var = 'ENCATT15'">Att: Subsys Type</xsl:when>
            <xsl:when test="$var = 'ENCATT16'">Att: Trx Class</xsl:when>
            <xsl:when test="$var = 'ENCATT17'">Att: Trx Name</xsl:when>
            <xsl:when test="$var = 'ENCATT18'">Att: User</xsl:when>
            <xsl:when test="$var = 'ENCATT19'">Att: Priority</xsl:when>
            <xsl:when test="$var = 'ENCATT20'">Att: Client IP Address</xsl:when>
            <xsl:when test="$var = 'ENCATT21'">Att: Client User ID</xsl:when>
            <xsl:when test="$var = 'ENCATT22'">Att: Client Transaction Name</xsl:when>
            <xsl:when test="$var = 'ENCATT23'">Att: Client Workstation/Host Name</xsl:when>
            <xsl:when test="$var = 'ENCATT24'">Att: Client Accounting Information</xsl:when>
            <!-- STOR -->
            <xsl:when test="$var = 'STRPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'STRPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'STRPCLA'">Type</xsl:when>
            <xsl:when test="$var = 'STRPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'STRPODEL'">% Delay ANY</xsl:when>
            <xsl:when test="$var = 'STR1SDEL'">% Delay COMM</xsl:when>
            <xsl:when test="$var = 'STR2SDEL'">% Delay LOCL</xsl:when>
            <xsl:when test="$var = 'STR3SDEL'">% Delay VIO</xsl:when>
            <xsl:when test="$var = 'STR4SDEL'">% Delay SWAP</xsl:when>
            <xsl:when test="$var = 'STR5SDEL'">% Delay OUTR</xsl:when>
            <xsl:when test="$var = 'STR6SDEL'">% Delay XMEM</xsl:when>
            <xsl:when test="$var = 'STR7SDEL'">% Delay HIPR</xsl:when>
            <xsl:when test="$var = 'STR8SDEL'">% Delay OTHR</xsl:when>
            <xsl:when test="$var = 'STRPACTV'">Active Frames</xsl:when>
            <xsl:when test="$var = 'STRPFIXD'">Fixed Frames</xsl:when>
            <xsl:when test="$var = 'STRPIDLE'">Idle</xsl:when>
            <xsl:when test="$var = 'STRPWSET'">Working Set</xsl:when>
            <xsl:when test="$var = 'STRPWSEX'">Working Set Expanded</xsl:when>
            <xsl:when test="$var = 'STRPJID'">JES ID</xsl:when>
            <!-- STORF -->
            <xsl:when test="$var = 'STFPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'STFPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'STFPCLA'">Type</xsl:when>
            <xsl:when test="$var = 'STFPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'STFPTOTL'">Total Frames</xsl:when>
            <xsl:when test="$var = 'STFPACTV'">Active Frames</xsl:when>
            <xsl:when test="$var = 'STFPIDLE'">Idle</xsl:when>
            <xsl:when test="$var = 'STFPWSET'">Working Set</xsl:when>
            <xsl:when test="$var = 'STFPFIXD'">Fixed Frames</xsl:when>
            <xsl:when test="$var = 'STFPDIV'">DIV Frames</xsl:when>
            <xsl:when test="$var = 'STFPAUXS'">AUX Frames</xsl:when>
            <xsl:when test="$var = 'STFPPGIN'">Page-in Rate</xsl:when>
            <xsl:when test="$var = 'STFPEXIN'">Page-in EXP</xsl:when>
            <xsl:when test="$var = 'STFPSPPI'">Page-in Shared Pages</xsl:when>
            <xsl:when test="$var = 'STFPTOTS'">Shared Page Views</xsl:when>
            <xsl:when test="$var = 'STFPSVIN'">Valid Shared Pages</xsl:when>
            <xsl:when test="$var = 'STFPSPVL'">Shared Valid Rate</xsl:when>
            <xsl:when test="$var = 'STFPLMO'">MemObjs 1 MB Fixed</xsl:when>
            <xsl:when test="$var = 'STFPLPR'">1 MB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STFPFREM'">Freemained Frames</xsl:when>
            <xsl:when test="$var = 'STFPGMO'">MemObjs 2 GB Fixed</xsl:when>
            <xsl:when test="$var = 'STFPGPR'">2 GB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STFDMNUK'">Dedicated Memory not used (4K)</xsl:when>
            <xsl:when test="$var = 'STFDMNUM'">Dedicated Memory not used (1M)</xsl:when>
            <xsl:when test="$var = 'STFDMNUG'">Dedicated Memory not used (2G)</xsl:when>
            <xsl:when test="$var = 'STFDMUPK'">Dedicated Memory used (4K)</xsl:when>
            <xsl:when test="$var = 'STFDMUPM'">Dedicated Memory used (1M pageable)</xsl:when>
            <xsl:when test="$var = 'STFDMUFM'">Dedicated Memory used (1M fixed)</xsl:when>
            <xsl:when test="$var = 'STFDMUFG'">Dedicated Memory used (2G fixed)</xsl:when>
            <xsl:when test="$var = 'STFDMMRG'">Minimum requested Dedicated Memory (2G)</xsl:when>
            <xsl:when test="$var = 'STFDMARG'">Requested Dedicated Memory (2G)</xsl:when>
            <xsl:when test="$var = 'STFDMAAG'">Assigned Dedicated Memory (2G)</xsl:when>
            <xsl:when test="$var = 'STFDMUSE'">Dedicated Memory usage indicator</xsl:when>
            <xsl:when test="$var = 'STFPJID'">JES ID</xsl:when>
            <!-- STORM -->
            <xsl:when test="$var = 'STMHSMO'">MemObjs Shared</xsl:when>
            <xsl:when test="$var = 'STMHCMO'">MemObjs Common</xsl:when>
            <xsl:when test="$var = 'STMHSFR'">Frames Shared</xsl:when>
            <xsl:when test="$var = 'STMHSSIZ'">Frames Shared Used %</xsl:when>
            <xsl:when test="$var = 'STMHCFR'">Frames Common</xsl:when>
            <xsl:when test="$var = 'STMHCSIZ'">Frames Common Used %</xsl:when>
            <xsl:when test="$var = 'STMHCFFR'">Frames Common Fixed</xsl:when>
            <xsl:when test="$var = 'STMHSASL'">AUX Slots Shared</xsl:when>
            <xsl:when test="$var = 'STMHCASL'">AUX Slots Common</xsl:when>
            <xsl:when test="$var = 'STMHLMO'">1 MB MemObjs Fixed</xsl:when>
            <xsl:when test="$var = 'STMHLCMO'">1 MB MemObjs Common</xsl:when>
            <xsl:when test="$var = 'STMHLCMU'">1 MB MemObjs Common Unowned</xsl:when>
            <xsl:when test="$var = 'STMHLSMO'">1 MB MemObjs Shared</xsl:when>
            <xsl:when test="$var = 'STMHLFR'">1 MB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STMHLSIZ'">1 MB LFAREA Maximum Used for 1 MB Fixed Pages %</xsl:when>
            <xsl:when test="$var = 'STMHLFF'">1 MB Frames Fixed Maximum</xsl:when>
            <xsl:when test="$var = 'STMHLF4K'">1 MB Frames Fixed for 4K Requests</xsl:when>
            <xsl:when test="$var = 'STMHLCFR'">1 MB Frames Fixed Common</xsl:when>
            <xsl:when test="$var = 'STMHLCPU'">1 MB Frames Fixed Common Unowned</xsl:when>
            <xsl:when test="$var = 'STMHFSIZ'">1 MB LFAREA Maximum Used %</xsl:when>
            <xsl:when test="$var = 'STMHLPF'">1 MB Frames Pageable Initial</xsl:when>
            <xsl:when test="$var = 'STMHLP4K'">1 MB Frames Pageable for 4K Requests</xsl:when>
            <xsl:when test="$var = 'STMHLFPF'">1 MB Frames Fixed for 1 MB Pageable Requests</xsl:when>
            <xsl:when test="$var = 'STMHLPFR'">1 MB Frames Pageable Failed</xsl:when>
            <xsl:when test="$var = 'STMHLPFC'">1 MB Frames Pageable Converted to 4K</xsl:when>
            <xsl:when test="$var = 'STMHPSIZ'">1 MB Frames Pageable Used %</xsl:when>
            <xsl:when test="$var = 'STMHUSIZ'">1 MB Frames Used %</xsl:when>
            <xsl:when test="$var = 'STMHLF'">1 MB Frames Total</xsl:when>
            <xsl:when test="$var = 'STMHGMO'">2 GB MemObjs Fixed</xsl:when>
            <xsl:when test="$var = 'STMHGFR'">2 GB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STMHGSIZ'">2 GB LFAREA Maximum Used %</xsl:when>
            <xsl:when test="$var = 'STMHGFF'">2 GB Frames Fixed Maximum</xsl:when>
            <xsl:when test="$var = 'STMHDMIA'">Dedicated Memory at IPL in 2G units (excl system)</xsl:when>
            <xsl:when test="$var = 'STMHDMOA'">Online Dedicated Memory in 2G units (incl system)</xsl:when>
            <xsl:when test="$var = 'STMHDMTA'">Total Dedicated Memory in 2G units (incl system)</xsl:when>
            <xsl:when test="$var = 'STMHDMFA'">Dedicated Memory available to address spaces in 2G units</xsl:when>
            <xsl:when test="$var = 'STMPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'STMPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'STMPCLA'">Class</xsl:when>
            <xsl:when test="$var = 'STMPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'STMPCLP'">Period</xsl:when>
            <xsl:when test="$var = 'STMPTMO'">MemObjs Total</xsl:when>
            <xsl:when test="$var = 'STMPCMO'">MemObjs Common</xsl:when>
            <xsl:when test="$var = 'STMPSMO'">MemObjs Shared</xsl:when>
            <xsl:when test="$var = 'STMPPMO'">MemObjs Private</xsl:when>
            <xsl:when test="$var = 'STMPLMO'">MemObjs 1 MB Fixed</xsl:when>
            <xsl:when test="$var = 'STMPLFR'">1 MB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STMPLFF'">1 MB Frames Page-Fixed</xsl:when>
            <xsl:when test="$var = 'STMPLPF'">1 MB Frames Pageable</xsl:when>
            <xsl:when test="$var = 'STMPVTB'">Bytes Total</xsl:when>
            <xsl:when test="$var = 'STMPCMB'">Bytes Common</xsl:when>
            <xsl:when test="$var = 'STMPVSB'">Bytes Shared</xsl:when>
            <xsl:when test="$var = 'STMPPMB'">Bytes Private</xsl:when>
            <xsl:when test="$var = 'STMPHCB'">Common HWM</xsl:when>
            <xsl:when test="$var = 'STMPLMB'">Memory Limit</xsl:when>
            <xsl:when test="$var = 'STMPLSMO'">MemObjs 1 MB Shared</xsl:when>
            <xsl:when test="$var = 'STMPHSB'">Shared HWM</xsl:when>
            <xsl:when test="$var = 'STMPGMO'">MemObjs 2 GB Fixed</xsl:when>
            <xsl:when test="$var = 'STMPGFR'">2 GB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'STMPJID'">JES ID</xsl:when>
            <!-- LOCKSP-->
            <xsl:when test="$var = 'LSPPRES'">Resource</xsl:when>
            <xsl:when test="$var = 'LSPPTYPE'">Type</xsl:when>
            <xsl:when test="$var = 'LSPPJT'">Job Name</xsl:when>
            <xsl:when test="$var = 'LSPPCPID'">CPU ID</xsl:when>
            <xsl:when test="$var = 'LSPPAC'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'LSPPRAD'">Address</xsl:when>
            <xsl:when test="$var = 'LSPPHELD'">% Held</xsl:when>
            <xsl:when test="$var = 'LSPPSPIN'">% Spin</xsl:when>
            <xsl:when test="$var = 'LSPPJID'">JES ID</xsl:when>
            <!-- LOCKSU-->
            <xsl:when test="$var = 'LSUPRES'">Resource</xsl:when>
            <xsl:when test="$var = 'LSUPTYPE'">Lock Type</xsl:when>
            <xsl:when test="$var = 'LSUPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'LSUPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'LSUPRAD'">Address</xsl:when>
            <xsl:when test="$var = 'LSUPHELD'">% Holding</xsl:when>
            <xsl:when test="$var = 'LSUPINTR'">% Interrupted</xsl:when>
            <xsl:when test="$var = 'LSUPDISP'">% Dispatchable</xsl:when>
            <xsl:when test="$var = 'LSUPSUSP'">% Suspended</xsl:when>
            <xsl:when test="$var = 'LSUPJID'">JES ID</xsl:when>
            <!-- CACHDET -->
            <xsl:when test="$var = 'CADHCDAT'">CDate</xsl:when>
            <xsl:when test="$var = 'CADHCTIM'">CTime</xsl:when>
            <xsl:when test="$var = 'CADHCRNG'">CRange</xsl:when>
            <xsl:when test="$var = 'CADPVOLU'">Volser</xsl:when>
            <xsl:when test="$var = 'CADPDVN5'">Device Number</xsl:when>
            <xsl:when test="$var = 'CADPSSID'">SSID</xsl:when>
            <xsl:when test="$var = 'CADPIOP'">I/O %</xsl:when>
            <xsl:when test="$var = 'CADPIO'">I/O Rate</xsl:when>
            <xsl:when test="$var = 'CADPHITP'">Hit %</xsl:when>
            <xsl:when test="$var = 'CADPREAD'">Cache Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADPDFW'">DFW Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADPCFW'">CFW Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADPTOT'">DASD I/O Rate</xsl:when>
            <xsl:when test="$var = 'CADPSTAG'">Stage I/O Rate</xsl:when>
            <xsl:when test="$var = 'CADPSEQ'">Seq Rate</xsl:when>
            <xsl:when test="$var = 'CADPASYN'">Async Rate</xsl:when>
            <xsl:when test="$var = 'CADVCACH'">Cache State</xsl:when>
            <xsl:when test="$var = 'CADVDFW'">DFW State</xsl:when>
            <xsl:when test="$var = 'CADVPIN'">Pinned State</xsl:when>
            <xsl:when test="$var = 'CADVNRRA'">Read Rate</xsl:when>
            <xsl:when test="$var = 'CADVNRHI'">Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVNRHP'">Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADVNWRA'">Write Rate</xsl:when>
            <xsl:when test="$var = 'CADVNWFA'">Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADVNWHI'">Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVNWHP'">Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADVNREP'">Read %</xsl:when>
            <xsl:when test="$var = 'CADVNTRA'">Tracks Rate</xsl:when>
            <xsl:when test="$var = 'CADVSRRA'">Sequential Read Rate</xsl:when>
            <xsl:when test="$var = 'CADVSRHI'">Sequential Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVSRHP'">Sequential Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADVSWRA'">Sequential Write Rate</xsl:when>
            <xsl:when test="$var = 'CADVSWFA'">Sequential Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADVSWHI'">Sequential Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVSWHP'">Sequential Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADVSREP'">Sequential Read %</xsl:when>
            <xsl:when test="$var = 'CADVSTRA'">Sequential Tracks Rate</xsl:when>
            <xsl:when test="$var = 'CADVCRRA'">CFW Read Rate</xsl:when>
            <xsl:when test="$var = 'CADVCRHI'">CFW Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVCRHP'">CFW Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADVCWRA'">CFW Write Rate</xsl:when>
            <xsl:when test="$var = 'CADVCWFA'">CFW Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADVCWHI'">CFW Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVCWHP'">CFW Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADVCREP'">CFW Read %</xsl:when>
            <xsl:when test="$var = 'CADVTRRA'">Total Read Rate</xsl:when>
            <xsl:when test="$var = 'CADVTRHI'">Total Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVTRHP'">Total Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADVTWRA'">Total Write Rate</xsl:when>
            <xsl:when test="$var = 'CADVTWFA'">Total Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADVTWHI'">Total Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADVTWHP'">Total Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADVTREP'">Total Read %</xsl:when>
            <xsl:when test="$var = 'CADVMMDB'">Del Op NVS</xsl:when>
            <xsl:when test="$var = 'CADVMNCI'">Non-cache ICL</xsl:when>
            <xsl:when test="$var = 'CADVMCWR'">CKD Write</xsl:when>
            <xsl:when test="$var = 'CADVMCRM'">Record Caching Read Miss</xsl:when>
            <xsl:when test="$var = 'CADVMMCB'">Del Op Cache</xsl:when>
            <xsl:when test="$var = 'CADVMNCB'">Non-cache Bypass</xsl:when>
            <xsl:when test="$var = 'CADVMCHI'">CKD Hits</xsl:when>
            <xsl:when test="$var = 'CADVMCWP'">Record Caching Write Prom</xsl:when>
            <xsl:when test="$var = 'CADVMMDI'">DFW Inhibit</xsl:when>
            <xsl:when test="$var = 'CADSNRRA'">SSID Normal Read Rate</xsl:when>
            <xsl:when test="$var = 'CADSNRHI'">SSID Normal Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSNRHP'">SSID Normal Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADSNWRA'">SSID Normal Write Rate</xsl:when>
            <xsl:when test="$var = 'CADSNWFA'">SSID Normal Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADSNWHI'">SSID Normal Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSNWHP'">SSID Normal Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADSNREP'">SSID Normal Read %</xsl:when>
            <xsl:when test="$var = 'CADSNTRA'">SSID Normal Tracks Rate</xsl:when>
            <xsl:when test="$var = 'CADSSRRA'">SSID Sequential Read Rate</xsl:when>
            <xsl:when test="$var = 'CADSSRHI'">SSID Sequential Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSSRHP'">SSID Sequential Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADSSWRA'">SSID Sequential Write Rate</xsl:when>
            <xsl:when test="$var = 'CADSSWFA'">SSID Sequential Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADSSWHI'">SSID Sequential Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSSWHP'">SSID Sequential Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADSSREP'">SSID Sequential Read %</xsl:when>
            <xsl:when test="$var = 'CADSSTRA'">SSID Sequential Tracks Rate</xsl:when>
            <xsl:when test="$var = 'CADSCRRA'">SSID CFW Read Rate</xsl:when>
            <xsl:when test="$var = 'CADSCRHI'">SSID CFW Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSCRHP'">SSID CFW Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADSCWRA'">SSID CFW Write Rate</xsl:when>
            <xsl:when test="$var = 'CADSCWFA'">SSID CFW Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADSCWHI'">SSID CFW Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSCWHP'">SSID CFW Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADSCREP'">SSID CFW Read %</xsl:when>
            <xsl:when test="$var = 'CADSTRRA'">SSID Total Read Rate</xsl:when>
            <xsl:when test="$var = 'CADSTRHI'">SSID Total Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSTRHP'">SSID Total Read Hit %</xsl:when>
            <xsl:when test="$var = 'CADSTWRA'">SSID Total Write Rate</xsl:when>
            <xsl:when test="$var = 'CADSTWFA'">SSID Total Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CADSTWHI'">SSID Total Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CADSTWHP'">SSID Total Write Hit %</xsl:when>
            <xsl:when test="$var = 'CADSTREP'">SSID Total Read %</xsl:when>
            <xsl:when test="$var = 'CADSMMDB'">SSID Del Op NVS</xsl:when>
            <xsl:when test="$var = 'CADSMNCI'">SSID Non-cache ICL</xsl:when>
            <xsl:when test="$var = 'CADSMCWR'">SSID CKD Write</xsl:when>
            <xsl:when test="$var = 'CADSMCRM'">SSID Record Caching Read Miss</xsl:when>
            <xsl:when test="$var = 'CADSMMCB'">SSID Del Op Cache</xsl:when>
            <xsl:when test="$var = 'CADSMNCB'">SSID Non-cache Bypass</xsl:when>
            <xsl:when test="$var = 'CADSMCHI'">SSID CKD Hits</xsl:when>
            <xsl:when test="$var = 'CADSMCWP'">SSID Record Caching Write Prom</xsl:when>
            <xsl:when test="$var = 'CADSMMDI'">SSID DFW Inhibit</xsl:when>
            <xsl:when test="$var = 'CADPDEVN'">4-Digit Device Number</xsl:when>
            <!-- CHANNEL -->
            <xsl:when test="$var = 'CHACPIVC'">CHPID</xsl:when>
            <xsl:when test="$var = 'CHACPTVC'">Type</xsl:when>
            <xsl:when test="$var = 'CHACSIVC'">Shared</xsl:when>
            <xsl:when test="$var = 'CHACPUVC'">LPAR Util %</xsl:when>
            <xsl:when test="$var = 'CHACTUVC'">Total Util %</xsl:when>
            <xsl:when test="$var = 'CHACTBVC'">Bus Util %</xsl:when>
            <xsl:when test="$var = 'CHACPRVC'">LPAR Read Bandwidth</xsl:when>
            <xsl:when test="$var = 'CHACTRVC'">Total Read Bandwidth</xsl:when>
            <xsl:when test="$var = 'CHACPWVC'">LPAR Write Bandwidth</xsl:when>
            <xsl:when test="$var = 'CHACTWVC'">Total Write Bandwidth</xsl:when>
            <xsl:when test="$var = 'CHACPMVC'">LPAR MSG Rate</xsl:when>
            <xsl:when test="$var = 'CHACPSVC'">LPAR MSG Size</xsl:when>
            <xsl:when test="$var = 'CHACSFVC'">LPAR Send Fail</xsl:when>
            <xsl:when test="$var = 'CHACPFVC'">LPAR Receive Fail</xsl:when>
            <xsl:when test="$var = 'CHACTMVC'">Total MSG Rate</xsl:when>
            <xsl:when test="$var = 'CHACTFVC'">Total Receive Fail</xsl:when>
            <xsl:when test="$var = 'CHACTSVC'">Total MSG Size</xsl:when>
            <xsl:when test="$var = 'CHACLTVC'">Line Type (DDS)</xsl:when>
            <xsl:when test="$var = 'CHACMGVC'">Channel Measurement Group</xsl:when>
            <xsl:when test="$var = 'CHACFRTE'">FICON Operation Rate</xsl:when>
            <xsl:when test="$var = 'CHACFACT'">FICON Active Operations</xsl:when>
            <xsl:when test="$var = 'CHACFDFR'">FICON Deferred Operation Rate</xsl:when>
            <xsl:when test="$var = 'CHACXRTE'">zHPF Operation Rate</xsl:when>
            <xsl:when test="$var = 'CHACXACT'">zHPF Active Operations</xsl:when>
            <xsl:when test="$var = 'CHACXDFR'">zHPF Deferred Operation Rate</xsl:when>
            <xsl:when test="$var = 'CHACNET1'">Physical Network ID Port 1</xsl:when>
            <xsl:when test="$var = 'CHACNET2'">Physical Network ID Port 2</xsl:when>
            <!-- DELAY -->
            <xsl:when test="$var = 'JDELDAN'">Job Name</xsl:when>
            <xsl:when test="$var = 'JDELASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'JDETYPX'">Job Class</xsl:when>
            <xsl:when test="$var = 'JDEPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'JDEPRPCL'">Report Class</xsl:when>
            <xsl:when test="$var = 'JDEGMIP'">Storage/CPU Critical</xsl:when>
            <xsl:when test="$var = 'JDELWFL'">% Workflow</xsl:when>
            <xsl:when test="$var = 'JDELUSG'">% Using</xsl:when>
            <xsl:when test="$var = 'JDELDEL'">% Delay</xsl:when>
            <xsl:when test="$var = 'JDELIDL'">% Idle</xsl:when>
            <xsl:when test="$var = 'JDELUKN'">% Unknown</xsl:when>
            <xsl:when test="$var = 'JDELPROC'">% Delay PROC</xsl:when>
            <xsl:when test="$var = 'JDELDEV'">% Delay DEV</xsl:when>
            <xsl:when test="$var = 'JDELSTOR'">% Delay STOR</xsl:when>
            <xsl:when test="$var = 'JDELSUBS'">% Delay SUBS</xsl:when>
            <xsl:when test="$var = 'JDELOPER'">% Delay OPER</xsl:when>
            <xsl:when test="$var = 'JDELENQ'">% Delay ENQ</xsl:when>
            <xsl:when test="$var = 'JDELJES'">% Delay JES</xsl:when>
            <xsl:when test="$var = 'JDELHSM'">% Delay HSM</xsl:when>
            <xsl:when test="$var = 'JDELXCF'">% Delay XCF</xsl:when>
            <xsl:when test="$var = 'JDELMNT'">% Delay Mount</xsl:when>
            <xsl:when test="$var = 'JDELMES'">% Delay MSG</xsl:when>
            <xsl:when test="$var = 'JDELQUI'">% Delay Quiesce</xsl:when>
            <xsl:when test="$var = 'JDELCAP'">% Delay Capping</xsl:when>
            <xsl:when test="$var = 'JDELCP'">% Delay CP</xsl:when>
            <xsl:when test="$var = 'JDELIFA'">% Delay AAP</xsl:when>
            <xsl:when test="$var = 'JDELSUP'">% Delay IIP</xsl:when>
            <xsl:when test="$var = 'JDELREAS'">Primary Reason for Delay</xsl:when>
            <xsl:when test="$var = 'JDELJID'">JES ID</xsl:when>
            <!-- USAGE -->
            <xsl:when test="$var = 'JUSPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'JUSPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'JUSPCLA'">Job Class</xsl:when>
            <xsl:when test="$var = 'JUSPCLAX'">Job Class Ext</xsl:when>
            <xsl:when test="$var = 'JUSPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'JUSPCLP'">Period</xsl:when>
            <xsl:when test="$var = 'JUSPDP'">Dispatching Priority</xsl:when>
            <xsl:when test="$var = 'JUSPTAT'">Transaction Active Time</xsl:when>
            <xsl:when test="$var = 'JUSPTRT'">Transaction Resident Time</xsl:when>
            <xsl:when test="$var = 'JUSPTCT'">Transaction Count</xsl:when>
            <xsl:when test="$var = 'JUSPFRT'">Total Frames</xsl:when>
            <xsl:when test="$var = 'JUSPFRXT'">Fixed Frames</xsl:when>
            <xsl:when test="$var = 'JUSPFRXH'">Fixed Frames High</xsl:when>
            <xsl:when test="$var = 'JUSPFRXA'">Fixed Frames Above</xsl:when>
            <xsl:when test="$var = 'JUSPFRXB'">Fixed Frames Below</xsl:when>
            <xsl:when test="$var = 'JUSPDCTT'">Total Device Connect Time</xsl:when>
            <xsl:when test="$var = 'JUSPDCTD'">Device Connect Time</xsl:when>
            <xsl:when test="$var = 'JUSPEXCT'">Total EXCP Count</xsl:when>
            <xsl:when test="$var = 'JUSPEXCD'">EXCP Count</xsl:when>
            <xsl:when test="$var = 'JUSPEXCR'">EXCP Rate</xsl:when>
            <xsl:when test="$var = 'JUSPCPUT'">Total CPU Time</xsl:when>
            <xsl:when test="$var = 'JUSPCPUD'">CPU Time</xsl:when>
            <xsl:when test="$var = 'JUSPTCBT'">Total TCB Time</xsl:when>
            <xsl:when test="$var = 'JUSPTCBD'">TCB Time</xsl:when>
            <xsl:when test="$var = 'JUSPQREQ'">QSCAN Requests</xsl:when>
            <xsl:when test="$var = 'JUSPQSPR'">Specific QSCAN Requests</xsl:when>
            <xsl:when test="$var = 'JUSPQRES'">QSCAN Resource Count</xsl:when>
            <xsl:when test="$var = 'JUSPQRSD'">QSCAN Resource Count Std.Dev.</xsl:when>
            <xsl:when test="$var = 'JUSPQTIM'">QSCAN Request Time</xsl:when>
            <xsl:when test="$var = 'JUSPQTSD'">QSCAN Request Time Std.Dev.</xsl:when>
            <xsl:when test="$var = 'JUSPRPB'">Recovery Process Boost</xsl:when>
            <xsl:when test="$var = 'JUSPJID'">JES ID</xsl:when>
            <!-- DEV -->
            <xsl:when test="$var = 'DEVPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'DEVPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'DEVPCLA'">Job Class</xsl:when>
            <xsl:when test="$var = 'DEVPODEL'">% Delay</xsl:when>
            <xsl:when test="$var = 'DEVPOUSE'">% Using</xsl:when>
            <xsl:when test="$var = 'DEVPCON'">% Connect</xsl:when>
            <xsl:when test="$var = 'DEVPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'DEV1SDEL'">% Delay (1)</xsl:when>
            <xsl:when test="$var = 'DEV1VOLU'">Volume Name (1)</xsl:when>
            <xsl:when test="$var = 'DEV2SDEL'">% Delay (2)</xsl:when>
            <xsl:when test="$var = 'DEV2VOLU'">Volume Name (2)</xsl:when>
            <xsl:when test="$var = 'DEV3SDEL'">% Delay (3)</xsl:when>
            <xsl:when test="$var = 'DEV3VOLU'">Volume Name (3)</xsl:when>
            <xsl:when test="$var = 'DEV4SDEL'">% Delay (4)</xsl:when>
            <xsl:when test="$var = 'DEV4VOLU'">Volume Name (4)</xsl:when>
            <xsl:when test="$var = 'DEVPJID'">JES ID</xsl:when>
            <!-- IOQUEUE-->
            <xsl:when test="$var = 'IOQCPIVC'">CHPID</xsl:when>
            <xsl:when test="$var = 'IOQPAAVC'">Path</xsl:when>
            <xsl:when test="$var = 'IOQDCMVC'">DCM</xsl:when>
            <xsl:when test="$var = 'IOQPCUVC'">Control Units</xsl:when>
            <xsl:when test="$var = 'IOQMMNVC'">DCM Min Used</xsl:when>
            <xsl:when test="$var = 'IOQMMXVC'">DCM Max Used</xsl:when>
            <xsl:when test="$var = 'IOQMDFVC'">DCM Defined</xsl:when>
            <xsl:when test="$var = 'IOQLCUVC'">LCU</xsl:when>
            <xsl:when test="$var = 'IOQCRTVC'">Contention Rate</xsl:when>
            <xsl:when test="$var = 'IOQDQLVC'">Delay Queue Length</xsl:when>
            <xsl:when test="$var = 'IOQCSSVC'">Avg CSS Time</xsl:when>
            <xsl:when test="$var = 'IOQCPTVC'">CHPID Taken</xsl:when>
            <xsl:when test="$var = 'IOQSPBVC'">% DP Busy</xsl:when>
            <xsl:when test="$var = 'IOQCUBVC'">% CU Busy</xsl:when>
            <xsl:when test="$var = 'IOQCBTVC'">Avg CU Busy Time</xsl:when>
            <xsl:when test="$var = 'IOQCMRVC'">Avg Cmd Response Time</xsl:when>
            <xsl:when test="$var = 'IOQPLTVC'">Line Type (DDS)</xsl:when>
            <!-- PCIE  -->
            <xsl:when test="$var = 'PCIEPFID'">Function Id</xsl:when>
            <xsl:when test="$var = 'PCIESTAT'">Function Status</xsl:when>
            <xsl:when test="$var = 'PCIEPCID'">Function CHID</xsl:when>
            <xsl:when test="$var = 'PCIEDEVT'">Function Type</xsl:when>
            <xsl:when test="$var = 'PCIEALLT'">% Alloc Time</xsl:when>
            <xsl:when test="$var = 'PCIEJOBN'">Job Name</xsl:when>
            <xsl:when test="$var = 'PCIEASID'">ASID</xsl:when>
            <xsl:when test="$var = 'PCIEDMAN'"># DMA AS</xsl:when>
            <xsl:when test="$var = 'PCIELOOP'">Load Operations Rate</xsl:when>
            <xsl:when test="$var = 'PCIESTOP'">Store Operations Rate</xsl:when>
            <xsl:when test="$var = 'PCIESBOP'">Store Block Operations Rate</xsl:when>
            <xsl:when test="$var = 'PCIERFOP'">Refresh PCI Translations Operations Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDMAR'">Xfer Read Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDMAW'">Xfer Write Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDPKR'">Received Packets Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDPKT'">Transmitted Packets Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDWUP'">Work Unit Rate</xsl:when>
            <xsl:when test="$var = 'PCIEDAUT'">Adapter Utilization</xsl:when>
            <xsl:when test="$var = 'PCIEADAT'">Allocation Date</xsl:when>
            <xsl:when test="$var = 'PCIEATIM'">Allocation Time</xsl:when>
            <xsl:when test="$var = 'PCIEFTYP'">HWA Type</xsl:when>
            <xsl:when test="$var = 'PCIEFBSY'">HWA Time % Busy</xsl:when>
            <xsl:when test="$var = 'PCIEFTR'">HWA Transfer Rate</xsl:when>
            <xsl:when test="$var = 'PCIEFRET'">Request Execution Time</xsl:when>
            <xsl:when test="$var = 'PCIEFRES'">Request Execution Time StdDev</xsl:when>
            <xsl:when test="$var = 'PCIEFRQT'">Request Queue Time</xsl:when>
            <xsl:when test="$var = 'PCIEFRQS'">Request Queue Time StdDev</xsl:when>
            <xsl:when test="$var = 'PCIEFRSZ'">Request Size</xsl:when>
            <xsl:when test="$var = 'PCIE1RRC'">Compression Request Rate</xsl:when>
            <xsl:when test="$var = 'PCIE1TPC'">Compression Throughput</xsl:when>
            <xsl:when test="$var = 'PCIE1RCC'">Compression Ratio</xsl:when>
            <xsl:when test="$var = 'PCIE1RRD'">Decompression Request Rate</xsl:when>
            <xsl:when test="$var = 'PCIE1TPD'">Decompression Throughput</xsl:when>
            <xsl:when test="$var = 'PCIE1RCD'">Decompression Ratio</xsl:when>
            <xsl:when test="$var = 'PCIE1BPS'">Buffer Pool Memory Size</xsl:when>
            <xsl:when test="$var = 'PCIE1BPU'">Buffer Pool % Utilization</xsl:when>
            <xsl:when test="$var = 'PCIENET1'">Physical Network ID Port 1</xsl:when>
            <xsl:when test="$var = 'PCIENET2'">Physical Network ID Port 2</xsl:when>
            <xsl:when test="$var = 'PCIEPOID'">Port ID</xsl:when>
            <xsl:when test="$var = 'PCIESERN'">Synchronous I/O Link Serial Number</xsl:when>
            <xsl:when test="$var = 'PCIETYMO'">Synchronous I/O Link Type-Model</xsl:when>
            <xsl:when test="$var = 'PCIELKID'">Synchronous I/O Link ID</xsl:when>
            <xsl:when test="$var = 'PCIETBPC'">Synchronous I/O Time % Busy (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIESRR'">Synchronous I/O Request Success %</xsl:when>
            <xsl:when test="$var = 'PCIESRRC'">Synchronous I/O Request Success % (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIERRT'">Synchronous I/O Request Rate</xsl:when>
            <xsl:when test="$var = 'PCIERRTC'">Synchronous I/O Request Rate (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIETRRC'">Synchronous I/O Transfer Read Rate (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIEXRR'">Synchronous I/O Transfer Read Ratio</xsl:when>
            <xsl:when test="$var = 'PCIEXRRC'">Synchronous I/O Transfer Read Ratio (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIETWRC'">Synchronous I/O Transfer Write Rate (CPC)</xsl:when>
            <xsl:when test="$var = 'PCIEXWR'">Synchronous I/O Transfer Write Ratio</xsl:when>
            <xsl:when test="$var = 'PCIEXWRC'">Synchronous I/O Transfer Write Ratio (CPC)</xsl:when>
            <!-- PROC -->
            <xsl:when test="$var = 'PRCPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'PRCPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'PRCPCLA'">Job Class</xsl:when>
            <xsl:when test="$var = 'PRCPCLAX'">Job Class Ext</xsl:when>
            <xsl:when test="$var = 'PRCPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'PRCPRPCL'">Report Class</xsl:when>
            <xsl:when test="$var = 'PRCPODEL'">Total % Delay</xsl:when>
            <xsl:when test="$var = 'PRCPOUSE'">Total % Using</xsl:when>
            <xsl:when test="$var = 'PRCPTYPE'">Processor Type</xsl:when>
            <xsl:when test="$var = 'PRCPTST'">Total Appl %</xsl:when>
            <xsl:when test="$var = 'PRCPCAP'">Capping Delay %</xsl:when>
            <xsl:when test="$var = 'PRCPETST'">Total EAppl %</xsl:when>
            <xsl:when test="$var = 'PRCPAPPL'">Appl % for Processor Type</xsl:when>
            <xsl:when test="$var = 'PRCPEAPP'">EAppl % for Processor Type</xsl:when>
            <xsl:when test="$var = 'PRCPTWFL'">% Workflow</xsl:when>
            <xsl:when test="$var = 'PRCPTDEL'">% Delay for Processor Type</xsl:when>
            <xsl:when test="$var = 'PRCPTUSE'">% Using for Processor Type</xsl:when>
            <xsl:when test="$var = 'PRCPAACP'">AAP on CP % Using</xsl:when>
            <xsl:when test="$var = 'PRCPIICP'">IIP on CP % Using</xsl:when>
            <xsl:when test="$var = 'PRC1SDEL'">1st Holding Job %</xsl:when>
            <xsl:when test="$var = 'PRC1JOBN'">1st Holding Job Name</xsl:when>
            <xsl:when test="$var = 'PRC2SDEL'">2nd Holding Job %</xsl:when>
            <xsl:when test="$var = 'PRC2JOBN'">2nd Holding Job Name</xsl:when>
            <xsl:when test="$var = 'PRC3SDEL'">3rd Holding Job %</xsl:when>
            <xsl:when test="$var = 'PRC3JOBN'">3rd Holding Job Name</xsl:when>
            <xsl:when test="$var = 'PRCTCPUT'">Total CPU Time in mSec</xsl:when>
            <xsl:when test="$var = 'PRCPJID'">JES ID</xsl:when>
            <!-- prev. releases -->
            <xsl:when test="$var = 'PRCPTUCP'">AAP/IIP on CP % Using</xsl:when>
            <xsl:when test="$var = 'PRCPVEC'">Vector Time Ratio</xsl:when>
            <xsl:when test="$var = 'PRCPCAPD'">Capping Delay %</xsl:when>
            <xsl:when test="$var = 'PRCPTCBT'">TCB %</xsl:when>
            <xsl:when test="$var = 'PRCPSRBT'">SRB %</xsl:when>
            <xsl:when test="$var = 'PRCPPCST'">P/C SRB %</xsl:when>
            <xsl:when test="$var = 'PRCPEPST'">P/C SRB + Enclave %</xsl:when>
            <xsl:when test="$var = 'PRCPIFAT'">AAP %</xsl:when>
            <xsl:when test="$var = 'PRCPCPT'">CP %</xsl:when>
            <xsl:when test="$var = 'PRCPIFCT'">AAP on CP %</xsl:when>
            <!-- PROCU -->
            <xsl:when test="$var = 'PRUPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'PRUPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'PRUPCLA'">Job Class</xsl:when>
            <xsl:when test="$var = 'PRUPCLAX'">Job Class Ext</xsl:when>
            <xsl:when test="$var = 'PRUPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'PRUPRPCL'">Report Class</xsl:when>
            <xsl:when test="$var = 'PRUPCLP'">Period</xsl:when>
            <xsl:when test="$var = 'PRUPCPT'">Total Time on CP %</xsl:when>
            <xsl:when test="$var = 'PRUPAACT'">AAP Time on CP %</xsl:when>
            <xsl:when test="$var = 'PRUPIICT'">IIP Time on CP %</xsl:when>
            <xsl:when test="$var = 'PRUPCPE'">CP EAppl %</xsl:when>
            <xsl:when test="$var = 'PRUPAAPE'">AAP EAppl %</xsl:when>
            <xsl:when test="$var = 'PRUPIIPE'">IIP EAppl %</xsl:when>
            <xsl:when test="$var = 'PRUPTOTC'">Total Appl %</xsl:when>
            <xsl:when test="$var = 'PRUPTOTE'">Total EAppl %</xsl:when>
            <xsl:when test="$var = 'PRUPTCB'">TCB Time %</xsl:when>
            <xsl:when test="$var = 'PRUPSRB'">SRB Time %</xsl:when>
            <xsl:when test="$var = 'PRUPPCS'">P/C SRB %</xsl:when>
            <xsl:when test="$var = 'PRUPEPS'">P/C SRB + Enclave %</xsl:when>
            <xsl:when test="$var = 'PRUTCPUT'">Total CPU Time in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTCPT'">Total Time on CP in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTAACT'">AAP Time on CP in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTIICT'">IIP Time on CP in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTCPE'">CP EAppl in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTAAPE'">AAP EAppl in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTIIPE'">IIP EAppl in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTTOTC'">Total Appl in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTTOTE'">Total EAppl in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTTCB'">TCB Time in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTSRB'">SRB Time in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTPCS'">P/C SRB in mSec</xsl:when>
            <xsl:when test="$var = 'PRUTEPS'">P/C SRB + Enclave in mSec</xsl:when>
            <xsl:when test="$var = 'PRUPRPB'">Recovery Process Boost</xsl:when>
            <xsl:when test="$var = 'PRUPJID'">JES ID</xsl:when>
            <!-- STORCR/STORC -->
            <xsl:when test="$report = 'STORCR'">
                <xsl:choose>
                    <!-- STORCR -->
                    <xsl:when test="$var = 'CSXNAME'">Job Name</xsl:when>
                    <xsl:when test="$var = 'CSXJESID'">JES ID</xsl:when>
                    <xsl:when test="$var = 'CSXTDATE'">Termination Date</xsl:when>
                    <xsl:when test="$var = 'CSXTTIME'">Termination Time</xsl:when>
                    <xsl:when test="$var = 'CSXACSA'">CSA not Released</xsl:when>
                    <xsl:when test="$var = 'CSXAECS'">ECSA not Released</xsl:when>
                    <xsl:when test="$var = 'CSXARUC'">RUCSA not Released</xsl:when>
                    <xsl:when test="$var = 'CSXAERU'">ERUCSA not Released</xsl:when>
                    <xsl:when test="$var = 'CSXASQA'">SQA not Released</xsl:when>
                    <xsl:when test="$var = 'CSXAESQ'">ESQA not Released</xsl:when>
                    <xsl:otherwise>STORCR Unknown Field</xsl:otherwise>
                </xsl:choose>
            </xsl:when>
            <xsl:when test="$report = 'STORC'">
                <xsl:choose>
                    <!-- STORC -->
                    <xsl:when test="$var = 'CSXNAME'">Job Name</xsl:when>
                    <xsl:when test="$var = 'CSXACT'">Act</xsl:when>
                    <xsl:when test="$var = 'CSXCLA'">Type</xsl:when>
                    <xsl:when test="$var = 'CSXSCN'">Service Class</xsl:when>
                    <xsl:when test="$var = 'CSXASID'">ASID (dec)</xsl:when>
                    <xsl:when test="$var = 'CSXTIME'">Elapsed Time</xsl:when>
                    <xsl:when test="$var = 'CSXPCSA'">% Used CSA</xsl:when>
                    <xsl:when test="$var = 'CSXPECS'">% Used ECSA</xsl:when>
                    <xsl:when test="$var = 'CSXPRUC'">% Used RUCSA</xsl:when>
                    <xsl:when test="$var = 'CSXPERU'">% Used ERUCSA</xsl:when>
                    <xsl:when test="$var = 'CSXPSQA'">% Used SQA</xsl:when>
                    <xsl:when test="$var = 'CSXPESQ'">% Used ESQA</xsl:when>
                    <xsl:when test="$var = 'CSXACSA'">CSA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXAECS'">ECSA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXARUC'">RUCSA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXAERU'">ERUCSA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXASQA'">SQA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXAESQ'">ESQA Amount Used</xsl:when>
                    <xsl:when test="$var = 'CSXJESID'">JES ID</xsl:when>
                    <!-- STORC Header variables -->
                    <xsl:when test="$var = 'CSUCSAIA'">CSA IPL Def</xsl:when>
                    <xsl:when test="$var = 'CSUECSIA'">ECSA IPL Def</xsl:when>
                    <xsl:when test="$var = 'CSUSQAIA'">SQA IPL Def</xsl:when>
                    <xsl:when test="$var = 'CSUESQIA'">ESQA IPL Def</xsl:when>
                    <xsl:when test="$var = 'CSUCSAPP'">CSA Peak %</xsl:when>
                    <xsl:when test="$var = 'CSUECSPP'">ECSA Peak %</xsl:when>
                    <xsl:when test="$var = 'CSUSQAPP'">SQA Peak %</xsl:when>
                    <xsl:when test="$var = 'CSUESQPP'">ESQA Peak %</xsl:when>
                    <xsl:when test="$var = 'CSUCSAPA'">CSA Peak Amount</xsl:when>
                    <xsl:when test="$var = 'CSUECSPA'">ECSA Peak Amount</xsl:when>
                    <xsl:when test="$var = 'CSUSQAPA'">SQA Peak Amount</xsl:when>
                    <xsl:when test="$var = 'CSUESQPA'">ESQA Peak Amount</xsl:when>
                    <xsl:when test="$var = 'CSUCSACP'">CSA-SQA Conv %</xsl:when>
                    <xsl:when test="$var = 'CSUECSCP'">CSA-SQA Conv %</xsl:when>
                    <xsl:when test="$var = 'CSUCSACA'">CSA-SQA Conv Amount</xsl:when>
                    <xsl:when test="$var = 'CSUECSCA'">ECSA-ESQA Conv Amount</xsl:when>
                    <xsl:when test="$var = 'CSUCSASP'">CSA Average %</xsl:when>
                    <xsl:when test="$var = 'CSUECSSP'">ECSA Average %</xsl:when>
                    <xsl:when test="$var = 'CSUSQASP'">SQA Average %</xsl:when>
                    <xsl:when test="$var = 'CSUESQSP'">ESQA Average %</xsl:when>
                    <xsl:when test="$var = 'CSUCSASA'">CSA Average</xsl:when>
                    <xsl:when test="$var = 'CSUECSSA'">ECSA Average</xsl:when>
                    <xsl:when test="$var = 'CSUSQASA'">SQA Average</xsl:when>
                    <xsl:when test="$var = 'CSUESQSA'">ESQA Average</xsl:when>
                    <xsl:when test="$var = 'CSUCSAAP'">CSA Avail %</xsl:when>
                    <xsl:when test="$var = 'CSUECSAP'">ECSA Avail %</xsl:when>
                    <xsl:when test="$var = 'CSUSQAAP'">SQA Avail %</xsl:when>
                    <xsl:when test="$var = 'CSUESQAP'">ESQA Avail %</xsl:when>
                    <xsl:when test="$var = 'CSUCSAAA'">CSA Available</xsl:when>
                    <xsl:when test="$var = 'CSUECSAA'">ECSA Available</xsl:when>
                    <xsl:when test="$var = 'CSUSQAAA'">SQA Available</xsl:when>
                    <xsl:when test="$var = 'CSUESQAA'">ESQA Available</xsl:when>
                    <xsl:when test="$var = 'CSUCSARE'">CSA+SQA Unallocated</xsl:when>
                    <xsl:when test="$var = 'CSURUCRE'">RUCSA Unallocated</xsl:when>
                    <xsl:when test="$var = 'CSUERCRE'">ERUCSA Unallocated</xsl:when>
                    <xsl:otherwise>
                        <xsl:text>STORC: </xsl:text>
                        <xsl:value-of select="$var"/>
                    </xsl:otherwise>
                </xsl:choose>
            </xsl:when>
            <!-- DEVR -->
            <xsl:when test="$var = 'DVRPVOLU'">Volume / Dev Number</xsl:when>
            <xsl:when test="$var = 'DVRPDVN5'">Device Number</xsl:when>
            <xsl:when test="$var = 'DVRPIDEN'">Dev Type / CU Type</xsl:when>
            <xsl:when test="$var = 'DVRPSTAT'">Status / PAV count</xsl:when>
            <xsl:when test="$var = 'DVRPEXP'"># Exposures</xsl:when>
            <xsl:when test="$var = 'DVRPKIND'">Device Type</xsl:when>
            <xsl:when test="$var = 'DVRPLCUN'">LCU Number</xsl:when>
            <xsl:when test="$var = 'DVRPACTV'">% Active</xsl:when>
            <xsl:when test="$var = 'DVRPCONN'">% Connect</xsl:when>
            <xsl:when test="$var = 'DVRPDISC'">% Disconnect</xsl:when>
            <xsl:when test="$var = 'DVRPPEND'">% Pending</xsl:when>
            <xsl:when test="$var = 'DVRPDLYR'">Pending Reason</xsl:when>
            <xsl:when test="$var = 'DVRPDLYP'">% Reason</xsl:when>
            <xsl:when test="$var = 'DVRACTRT'">Activity Rate</xsl:when>
            <xsl:when test="$var = 'DVRRESPT'">Response Time</xsl:when>
            <xsl:when test="$var = 'DVRIOSQT'">IOS Queue Time</xsl:when>
            <xsl:when test="$var = 'DVRPDVBT'">% Device Busy</xsl:when>
            <xsl:when test="$var = 'DVRPCMRT'">% CMR Time</xsl:when>
            <xsl:when test="$var = 'DVRPCUBT'">% Control Unit Busy</xsl:when>
            <xsl:when test="$var = 'DVRPSPBT'">% Switch Port Busy</xsl:when>
            <xsl:when test="$var = 'DVRRPTHP'">Response Time (ms)</xsl:when>
            <xsl:when test="$var = 'DVRIQTHP'">IOS Queue Time (ms)</xsl:when>
            <xsl:when test="$var = 'DVRIOINT'">IO Intensity</xsl:when>
            <xsl:when test="$var = 'DVRPJOBN'">Job Name</xsl:when>
            <xsl:when test="$var = 'DVRPASI'">ASID (dec)</xsl:when>
            <xsl:when test="$var = 'DVRPCLA'">Job Class</xsl:when>
            <xsl:when test="$var = 'DVRPDMN'">Domain</xsl:when>
            <xsl:when test="$var = 'DVRPPGN'">Performance Group</xsl:when>
            <xsl:when test="$var = 'DVRPSUSE'">% Using</xsl:when>
            <xsl:when test="$var = 'DVRPSDEL'">% Delay</xsl:when>
            <xsl:when test="$var = 'DVRPSVCL'">Service Class</xsl:when>
            <xsl:when test="$var = 'DVRPDEVN'">4-Digit Device Number</xsl:when>
            <!-- DSND -->
            <xsl:when test="$var = 'DNDPDSN'">Dataset Name</xsl:when>
            <xsl:when test="$var = 'DNDPVOLU'">Volume</xsl:when>
            <xsl:when test="$var = 'DNDPJOBN'">Job Name</xsl:when>
            <xsl:when test="$var = 'DNDPASID'">Job ASID</xsl:when>
            <xsl:when test="$var = 'DNDPDUSG'">Using %</xsl:when>
            <xsl:when test="$var = 'DNDPDDLY'">Delay %</xsl:when>
            <xsl:when test="$var = 'DNDPJID'">JES ID</xsl:when>
            <!-- CPC header variales -->
            <xsl:when test="$var = 'CPCHPNAM'">Partition Name</xsl:when>
            <xsl:when test="$var = 'CPCHMOD'">CPU Type</xsl:when>
            <xsl:when test="$var = 'CPCHMDL'">CPU Model</xsl:when>
            <xsl:when test="$var = 'CPCHBSTT'">Boost Type</xsl:when>
            <xsl:when test="$var = 'CPCHBSTC'">Boost Class</xsl:when>
            <xsl:when test="$var = 'CPCHCMSU'">CPC Capacity (MSU/h)</xsl:when>
            <xsl:when test="$var = 'CPCHWF'">Weight % of Max</xsl:when>
            <xsl:when test="$var = 'CPCHLMSU'">4h MSU Average</xsl:when>
            <xsl:when test="$var = 'CPCHIMSU'">Image Capacity</xsl:when>
            <xsl:when test="$var = 'CPCHCAP'">WLM Capping %</xsl:when>
            <xsl:when test="$var = 'CPCHLMAX'">4h MSU Maximum</xsl:when>
            <xsl:when test="$var = 'CPCHRMSU'">Proj Time until Capping</xsl:when>
            <xsl:when test="$var = 'CPCHCPU'">CPC sequence number</xsl:when>
            <xsl:when test="$var = 'CPCHCPCN'">CPC name</xsl:when>
            <xsl:when test="$var = 'CPCHCPNO'"># CP Processors</xsl:when>
            <xsl:when test="$var = 'CPCHICNO'"># ICF+IFL+AAP Processors</xsl:when>
            <xsl:when test="$var = 'CPCHIFAN'"># AAP Processors</xsl:when>
            <xsl:when test="$var = 'CPCHICFN'"># ICF Processors</xsl:when>
            <xsl:when test="$var = 'CPCHIFLN'"># IFL Processors</xsl:when>
            <xsl:when test="$var = 'CPCHSUPN'"># IIP Processors/Cores</xsl:when>
            <xsl:when test="$var = 'CPCHPANO'">Configured Partitions</xsl:when>
            <xsl:when test="$var = 'CPCHWAIT'">Wait Completion</xsl:when>
            <xsl:when test="$var = 'CPCHWMGT'">WLM LPAR management enabled</xsl:when>
            <xsl:when test="$var = 'CPCHVCPU'">Vary CPU management available</xsl:when>
            <xsl:when test="$var = 'CPCHPMSU'">% Capacity Used</xsl:when>
            <xsl:when test="$var = 'CPCHGNAM'">Capacity Group Name</xsl:when>
            <xsl:when test="$var = 'CPCHGLIM'">Capacity Group Limit</xsl:when>
            <xsl:when test="$var = 'CPCHGL4H'">Less than 4h in Capacity Group</xsl:when>
            <xsl:when test="$var = 'CPCHGAUN'">4h Unused Group Capacity Average</xsl:when>
            <xsl:when test="$var = 'CPCHRGRP'">Proj Time until Group Capping</xsl:when>
            <xsl:when test="$var = 'CPCHDEDC'"># Dedicated CPs</xsl:when>
            <xsl:when test="$var = 'CPCHDEDA'"># Dedicated AAPs</xsl:when>
            <xsl:when test="$var = 'CPCHDEDI'"># Dedicated IIPs</xsl:when>
            <xsl:when test="$var = 'CPCHSHRC'"># Shared physical CPs</xsl:when>
            <xsl:when test="$var = 'CPCHSHRA'"># Shared physical AAPs</xsl:when>
            <xsl:when test="$var = 'CPCHSHRI'"># Shared physical IIPs</xsl:when>
            <xsl:when test="$var = 'CPCHCUTL'">Physical Total % of shared CPs</xsl:when>
            <xsl:when test="$var = 'CPCHAUTL'">Physical Total % of shared AAPs</xsl:when>
            <xsl:when test="$var = 'CPCHLUTL'">Physical Total % of shared ICFs</xsl:when>
            <xsl:when test="$var = 'CPCHFUTL'">Physical Total % of shared IFLs</xsl:when>
            <xsl:when test="$var = 'CPCHUUTL'">Physical Total % of shared IIPs</xsl:when>
            <xsl:when test="$var = 'CPCHCCAI'">Capacity Adjustment Indicator</xsl:when>
            <xsl:when test="$var = 'CPCHCCCR'">Capacity Change Reason</xsl:when>
            <xsl:when test="$var = 'CPCHATDI'">Average Thread Density AAP</xsl:when>
            <xsl:when test="$var = 'CPCHATDS'">Average Thread Density IIP</xsl:when>
            <xsl:when test="$var = 'CPCHATD'">Average Thread Density CP</xsl:when>
            <xsl:when test="$var = 'CPCHCFI'">MT Capacity Factor AAP</xsl:when>
            <xsl:when test="$var = 'CPCHCFS'">MT Capacity Factor IIP</xsl:when>
            <xsl:when test="$var = 'CPCHCF'">MT Capacity Factor CP</xsl:when>
            <xsl:when test="$var = 'CPCHMCFI'">MT Max Capacity Factor AAP</xsl:when>
            <xsl:when test="$var = 'CPCHMCFS'">MT Max Capacity Factor IIP</xsl:when>
            <xsl:when test="$var = 'CPCHMCF'">MT Max Capacity Factor CP</xsl:when>
            <xsl:when test="$var = 'CPCHMTMI'">MT Mode AAP</xsl:when>
            <xsl:when test="$var = 'CPCHMTMS'">MT Mode IIP</xsl:when>
            <xsl:when test="$var = 'CPCHMTM'">MT Mode CP</xsl:when>
            <xsl:when test="$var = 'CPCHPRDI'">MT AAP Core Productivity</xsl:when>
            <xsl:when test="$var = 'CPCHPRDS'">MT IIP Core Productivity</xsl:when>
            <xsl:when test="$var = 'CPCHPRD'">MT CP Core Productivity</xsl:when>
            <xsl:when test="$var = 'CPCHAMSU'">Absolute MSU Capping</xsl:when>
            <!-- CPC -->
            <xsl:when test="$var = 'CPCPPNAM'">LPAR Name</xsl:when>
            <xsl:when test="$var = 'CPCPDMSU'">Defined MSU/h</xsl:when>
            <xsl:when test="$var = 'CPCPAMSU'">Actual MSU/h</xsl:when>
            <xsl:when test="$var = 'CPCPCAPI'">Initial Capping Option</xsl:when>
            <xsl:when test="$var = 'CPCPLPNO'"># Logical Processors/Cores Online</xsl:when>
            <xsl:when test="$var = 'CPCPLPND'"># Online Processors/Cores Shared</xsl:when>
            <xsl:when test="$var = 'CPCPDEDP'"># Online Processors/Cores Dedicated</xsl:when>
            <xsl:when test="$var = 'CPCPLEFU'">Logical Effective %</xsl:when>
            <xsl:when test="$var = 'CPCPLTOU'">Logical Total %</xsl:when>
            <xsl:when test="$var = 'CPCPPLMU'">LPAR Mgmt %</xsl:when>
            <xsl:when test="$var = 'CPCPPEFU'">Physical Effective %</xsl:when>
            <xsl:when test="$var = 'CPCPPTOU'">Physical Total %</xsl:when>
            <xsl:when test="$var = 'CPCPWGHT'">Current LPAR Weight</xsl:when>
            <xsl:when test="$var = 'CPCPLPSH'">Logical Processor Share %</xsl:when>
            <xsl:when test="$var = 'CPCPVCMH'">Hiper Dispatch: # High</xsl:when>
            <xsl:when test="$var = 'CPCPVCMM'">Hiper Dispatch: # Medium</xsl:when>
            <xsl:when test="$var = 'CPCPVCML'">Hiper Dispatch: # Low</xsl:when>
            <xsl:when test="$var = 'CPCPOSNM'">Operating System Name</xsl:when>
            <xsl:when test="$var = 'CPCPLPCN'">LPAR Cluster Name</xsl:when>
            <xsl:when test="$var = 'CPCPLCIW'">Initial Weight</xsl:when>
            <xsl:when test="$var = 'CPCPLCMW'">Min Weight</xsl:when>
            <xsl:when test="$var = 'CPCPLCXW'">Max Weight</xsl:when>
            <xsl:when test="$var = 'CPCPCGNM'">Group Capacity Name</xsl:when>
            <xsl:when test="$var = 'CPCPCGLT'">Group Capacity Limit</xsl:when>
            <xsl:when test="$var = 'CPCPCGEM'">Group Capacity Min Entitlement</xsl:when>
            <xsl:when test="$var = 'CPCPCGEX'">Group Capacity Max Entitlement</xsl:when>
            <xsl:when test="$var = 'CPCPCSMB'">Central Storage Online (MB)</xsl:when>
            <xsl:when test="$var = 'CPCPUPID'">User Partition ID</xsl:when>
            <xsl:when test="$var = 'CPCPIND'">Line Type</xsl:when>
            <xsl:when test="$var = 'CPCPBIIP'">zIIP Boost</xsl:when>
            <xsl:when test="$var = 'CPCPBSPD'">Speed Boost</xsl:when>
            <xsl:when test="$var = 'CPCPCAPD'">Capping Option</xsl:when>
            <xsl:when test="$var = 'CPCPHWCC'">Absolute Capping Limit (CPUs)</xsl:when>
            <xsl:when test="$var = 'CPCPHGNM'">Hardware Group Name</xsl:when>
            <xsl:when test="$var = 'CPCPHWGC'">Hardware Group Capping Limit (CPUs)</xsl:when>
            <!-- SYSSUM header variales -->
            <xsl:when test="$var = 'SUMSRVVC'">Service Definition</xsl:when>
            <xsl:when test="$var = 'SUMIDTVC'">Installed at</xsl:when>
            <xsl:when test="$var = 'SUMACPVC'">Active Policy</xsl:when>
            <xsl:when test="$var = 'SUMADTVC'">Activated at</xsl:when>
            <!-- SYSSUM -->
            <xsl:when test="$var = 'SUMGRP'">Name</xsl:when>
            <xsl:when test="$var = 'SUMTYP'">Type</xsl:when>
            <xsl:when test="$var = 'SUMRCTNT'">Tenant Report Class</xsl:when>
            <xsl:when test="$var = 'SUMIMP'">Imp</xsl:when>
            <xsl:when test="$var = 'SUMEVG'">Execution Vel Goal</xsl:when>
            <xsl:when test="$var = 'SUMEVA'">Execution Vel Actual</xsl:when>
            <xsl:when test="$var = 'SUMRTGTM'">RT Goal (ms)</xsl:when>
            <xsl:when test="$var = 'SUMRTGT'">RT Goal (sec)</xsl:when>
            <xsl:when test="$var = 'SUMRTGP'">RT Goal %</xsl:when>
            <xsl:when test="$var = 'SUMRTATM'">RT Actual (ms)</xsl:when>
            <xsl:when test="$var = 'SUMRTAT'">RT Actual (sec)</xsl:when>
            <xsl:when test="$var = 'SUMRTAP'">RT Actual %</xsl:when>
            <xsl:when test="$var = 'SUMPFID'">PI</xsl:when>
            <xsl:when test="$var = 'SUMTRAN'">Tran/sec</xsl:when>
            <xsl:when test="$var = 'SUMARTQM'">Queue Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTQ'">Queue Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTAM'">Active Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTA'">Active Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTTM'">Total Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTT'">Total Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTWM'">Wait Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTW'">Wait Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTCM'">Conv Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTC'">Conv Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTRM'">Res/Sys Affinity Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTR'">Res/Sys Affinity Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMARTIM'">Ineligible Queue Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMARTI'">Ineligible Queue Time (sec)</xsl:when>
            <xsl:when test="$var = 'SUMGOA'">Goaltype</xsl:when>
            <xsl:when test="$var = 'SUMDUR'">Duration</xsl:when>
            <xsl:when test="$var = 'SUMRES'">Resource Group</xsl:when>
            <xsl:when test="$var = 'SUMRGTYP'">Capacity Def</xsl:when>
            <xsl:when test="$var = 'SUMRGSPC'">Include Specialty Processors</xsl:when>
            <xsl:when test="$var = 'SUMSMI'">Capacity Min</xsl:when>
            <xsl:when test="$var = 'SUMSMA'">Capacity Max</xsl:when>
            <xsl:when test="$var = 'SUMSRA'">Capacity Actual</xsl:when>
            <xsl:when test="$var = 'SUMCRIT'">Storage/CPU Critical</xsl:when>
            <xsl:when test="$var = 'SUMHONP'">Honor Priority</xsl:when>
            <xsl:when test="$var = 'SUMMEMUS'">Memory Actual</xsl:when>
            <xsl:when test="$var = 'SUMMLIM'">Memory Limit</xsl:when>
            <xsl:when test="$var = 'SUMDDSIN'">Group Name (DDS)</xsl:when>
            <xsl:when test="$var = 'SUMDDSIT'">Type (DDS)</xsl:when>
            <xsl:when test="$var = 'SUMDDSIP'">Period (DDS)</xsl:when>
            <xsl:when test="$var = 'SUMEGRP'">Description</xsl:when>
            <xsl:when test="$var = 'SUMECTR'">Enclave Trans/sec</xsl:when>
            <xsl:when test="$var = 'SUMECTRE'">Avg Enclave Trans Execution Time (ms)</xsl:when>
            <xsl:when test="$var = 'SUMAIINF'">WLM Batch AI</xsl:when>
            <!-- SYSINFO header variables -->
            <xsl:when test="$var = 'SYSPARVC'">Partition Name</xsl:when>
            <xsl:when test="$var = 'SYSMODVC'">CPU Type</xsl:when>
            <xsl:when test="$var = 'SYSMDLVC'">CPU Model</xsl:when>
            <xsl:when test="$var = 'SYSTSVVC'">Appl%</xsl:when>
            <xsl:when test="$var = 'SYSIPVVC'">IPS Suffix</xsl:when>
            <xsl:when test="$var = 'SYSPOLVC'">WLM Policy Name</xsl:when>
            <xsl:when test="$var = 'SYSPRVVC'">CPs Online</xsl:when>
            <xsl:when test="$var = 'SYSCUVVC'">CPU Utilization % (CP)</xsl:when>
            <xsl:when test="$var = 'SYSTSEVC'">EAppl%</xsl:when>
            <xsl:when test="$var = 'SYSOPVVC'">OPT Suffix</xsl:when>
            <xsl:when test="$var = 'SYSPADVC'">Policy Date</xsl:when>
            <xsl:when test="$var = 'SYSPRIVC'">AAPs Online</xsl:when>
            <xsl:when test="$var = 'SYSLCPVC'">MVS Utilization % (CP)</xsl:when>
            <xsl:when test="$var = 'SYSAPIVC'">Appl% AAP</xsl:when>
            <xsl:when test="$var = 'SYSICVVC'">ICS Suffix</xsl:when>
            <xsl:when test="$var = 'SYSPATVC'">Policy Time</xsl:when>
            <xsl:when test="$var = 'SYSPRTVC'">IIPs Online</xsl:when>
            <xsl:when test="$var = 'SYSAPTVC'">Appl% IIP</xsl:when>
            <xsl:when test="$var = 'SYSVEPVC'">Vector Processors</xsl:when>
            <xsl:when test="$var = 'SYSAICVC'">Appl% AAP on CP</xsl:when>
            <xsl:when test="$var = 'SYSATCVC'">Appl% IIP on CP</xsl:when>
            <xsl:when test="$var = 'SYSLOAVG'">Load average</xsl:when>
            <xsl:when test="$var = 'SYSTCTVC'">Total CPU Time</xsl:when>
            <xsl:when test="$var = 'SYSUCTVC'">Uncaptured Time</xsl:when>
            <xsl:when test="$var = 'SYSCCTVC'">Captured Time</xsl:when>
            <xsl:when test="$var = 'SYSCULVC'">MVS Utilization % (CP)</xsl:when>
            <xsl:when test="$var = 'SYSCVAVC'">CPU Vary Activity found</xsl:when>
            <xsl:when test="$var = 'SYSCUAVC'">CPU Utilization % (AAP)</xsl:when>
            <xsl:when test="$var = 'SYSMUAVC'">MVS Utilization % (AAP)</xsl:when>
            <xsl:when test="$var = 'SYSCUIVC'">CPU Utilization % (IIP)</xsl:when>
            <xsl:when test="$var = 'SYSMUIVC'">MVS Utilization % (IIP)</xsl:when>
            <xsl:when test="$var = 'SYSAHPVC'">AAP Honor Priority</xsl:when>
            <xsl:when test="$var = 'SYSIHPVC'">IIP Honor Priority</xsl:when>
            <xsl:when test="$var = 'SYSPKCVC'">CPs Parked</xsl:when>
            <xsl:when test="$var = 'SYSPKAVC'">AAPs Parked</xsl:when>
            <xsl:when test="$var = 'SYSPKIVC'">IIPs Parked</xsl:when>
            <!-- SYSINFO -->
            <xsl:when test="$var = 'SYSNAMVC'">Group Name</xsl:when>
            <xsl:when test="$var = 'SYSTYPVC'">Type</xsl:when>
            <xsl:when test="$var = 'SYSWFLVC'">Workflow %</xsl:when>
            <xsl:when test="$var = 'SYSTUSVC'">Total Users</xsl:when>
            <xsl:when test="$var = 'SYSAUSVC'">Active Users</xsl:when>
            <xsl:when test="$var = 'SYSTRSVC'">Trans/sec</xsl:when>
            <xsl:when test="$var = 'SYSAFCVC'">Active Frames %</xsl:when>
            <xsl:when test="$var = 'SYSVECVC'">Vector Util %</xsl:when>
            <xsl:when test="$var = 'SYSAUPVC'">Avg # Using PROC</xsl:when>
            <xsl:when test="$var = 'SYSAUDVC'">Avg # Using DEV</xsl:when>
            <xsl:when test="$var = 'SYSADPVC'">Avg # Delayed PROC</xsl:when>
            <xsl:when test="$var = 'SYSADDVC'">Avg # Delayed DEV</xsl:when>
            <xsl:when test="$var = 'SYSADSVC'">Avg # Delayed STOR</xsl:when>
            <xsl:when test="$var = 'SYSADUVC'">Avg # Delayed SUBS</xsl:when>
            <xsl:when test="$var = 'SYSADOVC'">Avg # Delayed OPER</xsl:when>
            <xsl:when test="$var = 'SYSADEVC'">Avg # Delayed ENQ</xsl:when>
            <xsl:when test="$var = 'SYSADJVC'">Avg # Delayed JES</xsl:when>
            <xsl:when test="$var = 'SYSADHVC'">Avg # Delayed HSM</xsl:when>
            <xsl:when test="$var = 'SYSADXVC'">Avg # Delayed XCF</xsl:when>
            <xsl:when test="$var = 'SYSADNVC'">Avg # Delayed MOUNT</xsl:when>
            <xsl:when test="$var = 'SYSADMVC'">Avg # Delayed MSG</xsl:when>
            <xsl:when test="$var = 'SYSCPUVC'">Total CPU %</xsl:when>
            <xsl:when test="$var = 'SYSSRBVC'">Total SRB %</xsl:when>
            <xsl:when test="$var = 'SYSTCBVC'">Total TCB %</xsl:when>
            <xsl:when test="$var = 'SYSIFAVC'">Total AAP %</xsl:when>
            <xsl:when test="$var = 'SYSCPVC'">Total CP %</xsl:when>
            <xsl:when test="$var = 'SYSIFCVC'">Total AAP on CP %</xsl:when>
            <xsl:when test="$var = 'SYSRSPM'">Response Time (ms)</xsl:when>
            <xsl:when test="$var = 'SYSRSPVC'">Response Time (sec)</xsl:when>
            <xsl:when test="$var = 'SYSVELVC'">Execution Velocity</xsl:when>
            <xsl:when test="$var = 'SYSUGMVC'">% Using</xsl:when>
            <xsl:when test="$var = 'SYSUGPVC'">% Using PROC</xsl:when>
            <xsl:when test="$var = 'SYSUGDVC'">% Using DEV</xsl:when>
            <xsl:when test="$var = 'SYSWGDVC'">% Workflow DEV</xsl:when>
            <xsl:when test="$var = 'SYSWGPVC'">% Workflow PROC</xsl:when>
            <xsl:when test="$var = 'SYSDGMVC'">% Delay</xsl:when>
            <xsl:when test="$var = 'SYSUJMVC'">Avg # Using</xsl:when>
            <xsl:when test="$var = 'SYSDJMVC'">Avg # Delayed</xsl:when>
            <xsl:when test="$var = 'SYSDGEVC'">% Delay ENQ</xsl:when>
            <xsl:when test="$var = 'SYSDGHVC'">% Delay HSM</xsl:when>
            <xsl:when test="$var = 'SYSDGDVC'">% Delay DEV</xsl:when>
            <xsl:when test="$var = 'SYSDGJVC'">% Delay JES</xsl:when>
            <xsl:when test="$var = 'SYSDGOVC'">% Delay OPER</xsl:when>
            <xsl:when test="$var = 'SYSDGPVC'">% Delay PROC</xsl:when>
            <xsl:when test="$var = 'SYSDGSVC'">% Delay STOR</xsl:when>
            <xsl:when test="$var = 'SYSDGUVC'">% Delay SUBS</xsl:when>
            <xsl:when test="$var = 'SYSDGXVC'">% Delay XCF</xsl:when>
            <xsl:when test="$var = 'SYSDDSIN'">DDS Name</xsl:when>
            <xsl:when test="$var = 'SYSDDSIT'">DDS Type</xsl:when>
            <xsl:when test="$var = 'SYSDDSIP'">DDS Period</xsl:when>
            <xsl:when test="$var = 'SYSRCTNT'">Tenant Report Class</xsl:when>
            <xsl:when test="$var = 'SYSMEMUS'">Memory Actual</xsl:when>
            <xsl:when test="$var = 'SYSEAPVC'">EAppl %</xsl:when>
            <xsl:when test="$var = 'SYSLPIVC'">Local PI</xsl:when>
            <xsl:when test="$var = 'SYSSUPVC'">Total IIP %</xsl:when>
            <xsl:when test="$var = 'SYSSUCVC'">Total IIP on CP %</xsl:when>
            <xsl:when test="$var = 'SYSPDPVC'">CPU time at promoted DP</xsl:when>
            <xsl:when test="$var = 'SYSTODVC'">Total Delay Samples %</xsl:when>
            <xsl:when test="$var = 'SYSCPDVC'">CP Delay Samples %</xsl:when>
            <xsl:when test="$var = 'SYSECTVC'">Enclave Trans/sec</xsl:when>
            <xsl:when test="$var = 'SYSECEVC'">Avg Enclave Trans Execution Time (ms)</xsl:when>
            <xsl:when test="$var = 'SYSAPDVC'">AAP Delay Samples %</xsl:when>
            <xsl:when test="$var = 'SYSIPDVC'">IIP Delay Samples %</xsl:when>
            <xsl:when test="$var = 'SYSRGCVC'">CAPP Delay Samples %</xsl:when>
            <xsl:when test="$var = 'SYSDLYVC'">Total # of Delay Samples</xsl:when>
            <xsl:when test="$var = 'SYSCRIVC'">Storage/CPU Critical</xsl:when>
            <xsl:when test="$var = 'SYSAIIVC'">WLM Batch AI</xsl:when>
            <!-- CFACT -->
            <xsl:when test="$var = 'CFAPSTRU'">Structure Name</xsl:when>
            <xsl:when test="$var = 'CFAPTYPE'">Structure Type</xsl:when>
            <xsl:when test="$var = 'CFAPSTAT'">Structure Status</xsl:when>
            <xsl:when test="$var = 'CFAPSTEX'">Extended Structure Status</xsl:when>
            <xsl:when test="$var = 'CFAPENCR'">Encrypted</xsl:when>
            <xsl:when test="$var = 'CFAPSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'CFAPUTIP'">CF Utilization %</xsl:when>
            <xsl:when test="$var = 'CFAPSTEP'">Structure Execution %</xsl:when>
            <xsl:when test="$var = 'CFAPSYNR'">Sync Rate</xsl:when>
            <xsl:when test="$var = 'CFAPASS'">Sync Avg Service Time</xsl:when>
            <xsl:when test="$var = 'CFAPSYNC'">Sync Request Count</xsl:when>
            <xsl:when test="$var = 'CFAPASYR'">Async Rate</xsl:when>
            <xsl:when test="$var = 'CFAPAAS'">Async Avg Service Time</xsl:when>
            <xsl:when test="$var = 'CFAPASYC'">Async Request Count</xsl:when>
            <xsl:when test="$var = 'CFAPACHG'">Async Changed %</xsl:when>
            <xsl:when test="$var = 'CFAPADEL'">Async Delay %</xsl:when>
            <xsl:when test="$var = 'CFAPQRT'">Avg Queued Request Time</xsl:when>
            <xsl:when test="$var = 'CFAPMQRT'">Avg CF monopolization delay time</xsl:when>
            <xsl:when test="$var = 'CFAPCNVC'">Converted Request Count</xsl:when>
            <xsl:when test="$var = 'CFAPDELC'">Operation Count Delayed for Dump Serialization</xsl:when>
            <xsl:when test="$var = 'CFAPQUEC'">Queued Operation Count</xsl:when>
            <xsl:when test="$var = 'CFAPMONC'">Operation Count Delayed For CF Monopolization</xsl:when>
            <xsl:when test="$var = 'CFAPMUSR'">Maximum Number of Users</xsl:when>
            <xsl:when test="$var = 'CFAPTUSR'">Total Number of Users</xsl:when>
            <xsl:when test="$var = 'CFAPPUSR'">Number of Problem Users</xsl:when>
            <xsl:when test="$var = 'CFAPREBP'">Rebuild %</xsl:when>
            <xsl:when test="$var = 'CFAINAM'">CF Name</xsl:when>
            <xsl:when test="$var = 'CFAISTRU'">Structure Name</xsl:when>
            <xsl:when test="$var = 'CFAITYPE'">Structure Type</xsl:when>
            <xsl:when test="$var = 'CFAICNAM'">Connection Name</xsl:when>
            <xsl:when test="$var = 'CFAICJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'CFAICSTA'">Status</xsl:when>
            <xsl:when test="$var = 'CFAICPRB'">Problem Status</xsl:when>
            <xsl:when test="$var = 'CFAICASI'">ASID</xsl:when>
            <xsl:when test="$var = 'CFAICLVL'">Connect CF Level</xsl:when>
            <xsl:when test="$var = 'CFAICREB'">User-Managed Rebuild Allowed</xsl:when>
            <xsl:when test="$var = 'CFAICDRB'">User-Managed Rebuild with Duplexing Allowed</xsl:when>
            <xsl:when test="$var = 'CFAICALT'">Alter Allowed</xsl:when>
            <xsl:when test="$var = 'CFAICAUT'">System-Managed Processes Allowed</xsl:when>
            <xsl:when test="$var = 'CFAICSUS'">Suspension of Work Tolerated</xsl:when>
            <xsl:when test="$var = 'CFAISTRS'">Structure Size</xsl:when>
            <xsl:when test="$var = 'CFAISTRP'">Structure Size %</xsl:when>
            <xsl:when test="$var = 'CFAISTUP'">Utilized Structure Storage %</xsl:when>
            <xsl:when test="$var = 'CFAISTRC'">Structure Storage Class</xsl:when>
            <xsl:when test="$var = 'CFAISTRM'">Min. Structure Size</xsl:when>
            <xsl:when test="$var = 'CFAISTRX'">Max. Structure Size</xsl:when>
            <xsl:when test="$var = 'CFAIDTS'">Dump Table Size</xsl:when>
            <xsl:when test="$var = 'CFAILDES'">Data Elements Size (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAILDLS'">Data List Entry Size (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAILEL'">Total List Entries (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAILEM'">Actual List Entries (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAIMAE'">Total Data Elements (LIST/CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAICUE'">Actual Data Elements (LIST/CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAILTL'">Total Lock Entries (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAILTM'">Actual Lock Entries (LIST/LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAIDES'">Data Element Size (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIDEN'">Total Dir Entries (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIDEC'">Actual Dir Entries (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAICEN'">Changed Dir Entries (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIDEL'">Total Data Elements (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIDAC'">Actual Data Elements (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAICEL'">Changed Data Elements (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAICONT'">Contention % (LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAIFCON'">False Contention % (LOCK)</xsl:when>
            <xsl:when test="$var = 'CFAIREQR'">Request Rate</xsl:when>
            <xsl:when test="$var = 'CFAIREAR'">Read Rate (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIWRIR'">Write Rate (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAICAOR'">Castout Rate (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIXIR'">XI Rate (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIDER'">Directory Reclaims (CACHE)</xsl:when>
            <xsl:when test="$var = 'CFAIFCCL'">First Castout Class</xsl:when>
            <xsl:when test="$var = 'CFAILCCL'">Last Castout Class</xsl:when>
            <xsl:when test="$var = 'CFAIPREF'">Allocation Preference List</xsl:when>
            <xsl:when test="$var = 'CFAIEXCL'">Exclusion Preference List</xsl:when>
            <xsl:when test="$var = 'CFAISPCF'">% of CF storage</xsl:when>
            <xsl:when test="$var = 'CFAISCM'">SCM data available</xsl:when>
            <xsl:when test="$var = 'CFAISAUM'">Estimated Max. of Augmented Space</xsl:when>
            <xsl:when test="$var = 'CFAISAUP'">% Augmented Space Used</xsl:when>
            <xsl:when test="$var = 'CFAISSCM'">Maximum of SCM Space</xsl:when>
            <xsl:when test="$var = 'CFAISSCP'">% SCM Space Used</xsl:when>
            <xsl:when test="$var = 'CFAISLTM'">Estimated Max. of SCM List Entries</xsl:when>
            <xsl:when test="$var = 'CFAISLTC'">Actual SCM List Entries</xsl:when>
            <xsl:when test="$var = 'CFAISLMM'">Estimated Max. of SCM List Elements</xsl:when>
            <xsl:when test="$var = 'CFAISLMC'">Actual SCM List Elements</xsl:when>
            <xsl:when test="$var = 'CFAISALG'">SCM Algorithm Type</xsl:when>
            <!-- CFOVER -->
            <xsl:when test="$var = 'CFOHPNAM'">Policy Name</xsl:when>
            <xsl:when test="$var = 'CFOHPACD'">Policy Activation Date</xsl:when>
            <xsl:when test="$var = 'CFOHPACT'">Policy Activation Time</xsl:when>
            <xsl:when test="$var = 'CFOHPREF'">Policy Reformat Required</xsl:when>
            <xsl:when test="$var = 'CFOPNAM'">CF Name</xsl:when>
            <xsl:when test="$var = 'CFOPMOD'">Model</xsl:when>
            <xsl:when test="$var = 'CFOPVER'">Version</xsl:when>
            <xsl:when test="$var = 'CFOPLVL'">CF Level</xsl:when>
            <xsl:when test="$var = 'CFOPDYND'">CF Dynamic Dispatching</xsl:when>
            <xsl:when test="$var = 'CFOPSTAT'">Status of CF</xsl:when>
            <xsl:when test="$var = 'CFOPVOL'">CF Storage Volatile</xsl:when>
            <xsl:when test="$var = 'CFOPUTIP'">Proc Util %</xsl:when>
            <xsl:when test="$var = 'CFOPDEF'">Processors defined</xsl:when>
            <xsl:when test="$var = 'CFOPPDED'">Number of Dedicated Processors</xsl:when>
            <xsl:when test="$var = 'CFOPPSHR'">Number of Shared Processors</xsl:when>
            <xsl:when test="$var = 'CFOPPWGT'">Average Weighting of Shared Processors</xsl:when>
            <xsl:when test="$var = 'CFOPEFF'">Processors effective</xsl:when>
            <xsl:when test="$var = 'CFOPREQR'">Request Rate</xsl:when>
            <xsl:when test="$var = 'CFOPTSD'">Storage Size (in Bytes)</xsl:when>
            <xsl:when test="$var = 'CFOPTSF'">Storage Avail (in Bytes)</xsl:when>
            <xsl:when test="$var = 'CFOPUTIS'">Utilized Storage %</xsl:when>
            <xsl:when test="$var = 'CFOPTCS'">Total Control Space</xsl:when>
            <xsl:when test="$var = 'CFOPFCS'">Free Control Space</xsl:when>
            <xsl:when test="$var = 'CFOPDTS'">Dump Table Space</xsl:when>
            <xsl:when test="$var = 'CFOPDTUS'">Dump Table in Use</xsl:when>
            <xsl:when test="$var = 'CFOPSYSC'">Connected MVS System Count</xsl:when>
            <xsl:when test="$var = 'CFOPSTCI'">Structure Count in Policy</xsl:when>
            <xsl:when test="$var = 'CFOPSTCO'">Structure Count out Policy</xsl:when>
            <xsl:when test="$var = 'CFOPMNT'">Maintenance Mode</xsl:when>
            <xsl:when test="$var = 'CFOPRCV'">Recovery Mode</xsl:when>
            <xsl:when test="$var = 'CFOPSCMS'">Storage Class Memory Size</xsl:when>
            <xsl:when test="$var = 'CFOPSCMA'">Storage Class Memory Available</xsl:when>
            <xsl:when test="$var = 'CFOPSCMU'">% Utilized Storage Class Memory</xsl:when>
            <xsl:when test="$var = 'CFOPAUGS'">Augmented Space Maximum</xsl:when>
            <xsl:when test="$var = 'CFOPAUGA'">Augmented Space Available</xsl:when>
            <xsl:when test="$var = 'CFOPAUGU'">% Utilized Augmented Space</xsl:when>
            <xsl:when test="$var = 'CFOPSMSC'">Sum of maximum Storage Class Memory</xsl:when>
            <!-- CFSYS -->
            <xsl:when test="$var = 'CFSPNAM'">CF Name</xsl:when>
            <xsl:when test="$var = 'CFSPSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'CFSPSDEL'">Subchannel Delay %</xsl:when>
            <xsl:when test="$var = 'CFSPSBSP'">Subchannel Busy %</xsl:when>
            <xsl:when test="$var = 'CFSPPTHA'">Paths Available</xsl:when>
            <xsl:when test="$var = 'CFSPPDEL'">Paths Delay %</xsl:when>
            <xsl:when test="$var = 'CFSPSYNR'">Sync Rate</xsl:when>
            <xsl:when test="$var = 'CFSPASS'">Sync Avg Service Time</xsl:when>
            <xsl:when test="$var = 'CFSPASYR'">Async Rate</xsl:when>
            <xsl:when test="$var = 'CFSPAAS'">Async Avg Service Time</xsl:when>
            <xsl:when test="$var = 'CFSPACHG'">Async Changed %</xsl:when>
            <xsl:when test="$var = 'CFSPADEL'">Async Delay %</xsl:when>
            <xsl:when test="$var = 'CFSPREQC'">Total Request Count</xsl:when>
            <xsl:when test="$var = 'CFSPSYNC'">Sync Request Count</xsl:when>
            <xsl:when test="$var = 'CFSPASYC'">Async Request Count</xsl:when>
            <xsl:when test="$var = 'CFSPCNVC'">Sync to Async Conversion Count</xsl:when>
            <xsl:when test="$var = 'CFSPSYNP'">Sync Request %</xsl:when>
            <xsl:when test="$var = 'CFSPASYP'">Async Request %</xsl:when>
            <xsl:when test="$var = 'CFSPSOPD'">Avg Synchronous Operation Delay</xsl:when>
            <xsl:when test="$var = 'CFSPFOPT'">Avg Failed Operation Time</xsl:when>
            <xsl:when test="$var = 'CFSINAM'">CF Name</xsl:when>
            <xsl:when test="$var = 'CFSISCG'">Subchannels Generated</xsl:when>
            <xsl:when test="$var = 'CFSISCU'">Subchannels in Use</xsl:when>
            <xsl:when test="$var = 'CFSISCL'">Subchannels Max</xsl:when>
            <xsl:when test="$var = 'CFSICPI1'">Path1 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT1'">Path1 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO1'">Path1 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD1'">Path1 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL1'">Path1 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY1'">Path1 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA1'">Path1 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP1'">Path1 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP1'">Path1 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI2'">Path2 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT2'">Path2 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO2'">Path2 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD2'">Path2 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL2'">Path2 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY2'">Path2 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA2'">Path2 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP2'">Path2 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP2'">Path2 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI3'">Path3 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT3'">Path3 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO3'">Path3 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD3'">Path3 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL3'">Path3 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY3'">Path3 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA3'">Path3 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP3'">Path3 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP3'">Path3 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI4'">Path4 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT4'">Path4 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO4'">Path4 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD4'">Path4 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL4'">Path4 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY4'">Path4 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA4'">Path4 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP4'">Path4 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP4'">Path4 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI5'">Path5 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT5'">Path5 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO5'">Path5 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD5'">Path5 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL5'">Path5 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY5'">Path5 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA5'">Path5 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP5'">Path5 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP5'">Path5 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI6'">Path6 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT6'">Path6 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO6'">Path6 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD6'">Path6 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL6'">Path6 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY6'">Path6 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA6'">Path6 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP6'">Path6 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP6'">Path6 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI7'">Path7 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT7'">Path7 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO7'">Path7 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD7'">Path7 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL7'">Path7 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY7'">Path7 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA7'">Path7 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP7'">Path7 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP7'">Path7 IOP IDs</xsl:when>
            <xsl:when test="$var = 'CFSICPI8'">Path8 ID</xsl:when>
            <xsl:when test="$var = 'CFSICPT8'">Path8 Type</xsl:when>
            <xsl:when test="$var = 'CFSICPO8'">Path8 Operation Mode</xsl:when>
            <xsl:when test="$var = 'CFSICPD8'">Path8 Degraded</xsl:when>
            <xsl:when test="$var = 'CFSICPL8'">Path8 Distance</xsl:when>
            <xsl:when test="$var = 'CFSIPHY8'">Path8 CHID</xsl:when>
            <xsl:when test="$var = 'CFSIHCA8'">Path8 Adapter ID</xsl:when>
            <xsl:when test="$var = 'CFSIHCP8'">Path8 Adapter Port Number</xsl:when>
            <xsl:when test="$var = 'CFSIIOP8'">Path8 IOP IDs</xsl:when>
            <!-- XCFGROUP -->
            <xsl:when test="$var = 'XGRGRP'">Group Name</xsl:when>
            <xsl:when test="$var = 'XGRMEM'">Member Name</xsl:when>
            <xsl:when test="$var = 'XGRSTSH'">Status (short)</xsl:when>
            <xsl:when test="$var = 'XGRSTAT'">Status</xsl:when>
            <xsl:when test="$var = 'XGRINTV'">Status Checking Interval</xsl:when>
            <xsl:when test="$var = 'XGRSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'XGRJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'XGRROUT'">Outbound Requests</xsl:when>
            <xsl:when test="$var = 'XGRRIN'">Inbound Requests</xsl:when>
            <xsl:when test="$var = 'XGRLIND'">Line Type</xsl:when>
            <!-- XCFPATH -->
            <xsl:when test="$var = 'XPASYSP'">Systems</xsl:when>
            <xsl:when test="$var = 'XPASYSO'">System(1)</xsl:when>
            <xsl:when test="$var = 'XPATYPE'">Path Type</xsl:when>
            <xsl:when test="$var = 'XPAPATH'">Structure or CTC Devices</xsl:when>
            <xsl:when test="$var = 'XPASYSD'">System(2)</xsl:when>
            <xsl:when test="$var = 'XPATCN'">Transport Class</xsl:when>
            <xsl:when test="$var = 'XPASTSH'">Status (short)</xsl:when>
            <xsl:when test="$var = 'XPASTAT'">Status</xsl:when>
            <xsl:when test="$var = 'XPARETP'">Retry %</xsl:when>
            <xsl:when test="$var = 'XPARETL'">Retry Limit</xsl:when>
            <xsl:when test="$var = 'XPAMSGL'">Message Limit</xsl:when>
            <xsl:when test="$var = 'XPASENT'">Signals Sent</xsl:when>
            <xsl:when test="$var = 'XPABUSY'">Times Path Busy</xsl:when>
            <xsl:when test="$var = 'XPAPEND'">Signals Pending</xsl:when>
            <xsl:when test="$var = 'XPASUSE'">Storage in Use</xsl:when>
            <xsl:when test="$var = 'XPARCNT'">Restart Count</xsl:when>
            <xsl:when test="$var = 'XPARECV'">Signals Received</xsl:when>
            <xsl:when test="$var = 'XPABUNA'">Times Buffer Unavailable</xsl:when>
            <xsl:when test="$var = 'XPAIOXT'">I/O Transfer Time</xsl:when>
            <xsl:when test="$var = 'XPALIND'">Line Type</xsl:when>
            <!-- XCFSYS -->
            <xsl:when test="$var = 'XSYSYSP'">Systems</xsl:when>
            <xsl:when test="$var = 'XSYSYSO'">System(1)</xsl:when>
            <xsl:when test="$var = 'XSYSYSD'">System(2)</xsl:when>
            <xsl:when test="$var = 'XSYTCN'">Transport Class</xsl:when>
            <xsl:when test="$var = 'XSYSENT'">Signals Sent</xsl:when>
            <xsl:when test="$var = 'XSYRECV'">Signals Received</xsl:when>
            <xsl:when test="$var = 'XSYPUNA'">Times Path Unavailable</xsl:when>
            <xsl:when test="$var = 'XSYBUNA'">Times Buffer Unavailable</xsl:when>
            <xsl:when test="$var = 'XSYBLEN'">Buffer Length</xsl:when>
            <xsl:when test="$var = 'XSYFIT'">Fit %</xsl:when>
            <xsl:when test="$var = 'XSYSML'">Smaller %</xsl:when>
            <xsl:when test="$var = 'XSYLAR'">Larger %</xsl:when>
            <xsl:when test="$var = 'XSYDEG'">Degraded %</xsl:when>
            <xsl:when test="$var = 'XSYLIND'">Direction</xsl:when>
            <!-- XCFOVW -->
            <xsl:when test="$var = 'XSOSYS'">System Name</xsl:when>
            <xsl:when test="$var = 'XSOSID'">SMF ID</xsl:when>
            <xsl:when test="$var = 'XSOPART'">Partition Name</xsl:when>
            <xsl:when test="$var = 'XSOREL'">System Level</xsl:when>
            <xsl:when test="$var = 'XSOINTM'">Monitoring Interval</xsl:when>
            <xsl:when test="$var = 'XSOINTO'">Operator Interval</xsl:when>
            <xsl:when test="$var = 'XSOSTAT'">Status</xsl:when>
            <xsl:when test="$var = 'XSORMFM'">RMF Master</xsl:when>
            <!-- STORS and STORR header variables -->
            <xsl:when test="$var = 'SRHRHUIC'">System UIC</xsl:when>
            <xsl:when test="$var = 'SRHRSTOR'">Real Storage Frames</xsl:when>
            <xsl:when test="$var = 'SRHRNUC'">% Nucleus Frames</xsl:when>
            <xsl:when test="$var = 'SRHRSQA'">% SQA Frames</xsl:when>
            <xsl:when test="$var = 'SRHRCSA'">% CSA Frames</xsl:when>
            <xsl:when test="$var = 'SRHRLPA'">% LPA Frames</xsl:when>
            <xsl:when test="$var = 'SRHRACTV'">% Active Frames</xsl:when>
            <xsl:when test="$var = 'SRHRIDLE'">% Idle Frames</xsl:when>
            <xsl:when test="$var = 'SRHRAVAL'">% Available Frames</xsl:when>
            <xsl:when test="$var = 'SRHRSHR'">% Shared Frames</xsl:when>
            <xsl:when test="$var = 'SRHPGINR'">Pagein Rate</xsl:when>
            <xsl:when test="$var = 'SRHFREEF'"># Frames Available</xsl:when>
            <xsl:when test="$var = 'SRHFREES'"># Slots Available</xsl:when>
            <xsl:when test="$var = 'SRHFREEV'"># Frames and Slots Available</xsl:when>
            <xsl:when test="$var = 'SRHLMO'">1 MB MemObjs Fixed</xsl:when>
            <xsl:when test="$var = 'SRHLPR'">1 MB Frames Fixed</xsl:when>
            <xsl:when test="$var = 'SRHGMO'">2 GB MemObjs Fixed</xsl:when>
            <xsl:when test="$var = 'SRHGPR'">2 GB Frames Fixed</xsl:when>
            <!-- additional STORR header variables -->
            <xsl:when test="$var = 'SRHEMAGE'">Migration Age</xsl:when>
            <xsl:when test="$var = 'SRHESTOR'">Expanded Storage Frames</xsl:when>
            <xsl:when test="$var = 'SRHESQA'">% SQA Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHECSA'">% CSA Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHELPA'">% LPA Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHEACTV'">% Active Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHEIDLE'">% Idle Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHEAVAL'">% Available Frames Expanded</xsl:when>
            <xsl:when test="$var = 'SRHSQAOV'">SQA Overflow</xsl:when>
            <xsl:when test="$var = 'SRHLOAL'">User Region Pages Below 16 MB</xsl:when>
            <xsl:when test="$var = 'SRHHIAL'">LSQA/SWA Pages Below 16 MB</xsl:when>
            <xsl:when test="$var = 'SRHELOAL'">User Region Pages Above 16 MB</xsl:when>
            <xsl:when test="$var = 'SRHEHIAL'">LSQA/SWA Pages Above 16 MB</xsl:when>
            <xsl:when test="$var = 'SRHRTFIX'"># Fixed Frames</xsl:when>
            <xsl:when test="$var = 'SRHRBFIX'"># Fixed Frames Below 16 MB</xsl:when>
            <xsl:when test="$var = 'SRHRFREM'">Freemained Frames</xsl:when>
            <!-- STORS -->
            <xsl:when test="$var = 'SRSPDMPG'">Group Name</xsl:when>
            <xsl:when test="$var = 'SRSPSVCL'">Group Name</xsl:when>
            <xsl:when test="$var = 'SRSPGTYP'">Group Type</xsl:when>
            <xsl:when test="$var = 'SRSRCTNT'">Tenant Report Class</xsl:when>
            <xsl:when test="$var = 'SRSPTOTU'">Total Users</xsl:when>
            <xsl:when test="$var = 'SRSPACTU'">Active Users</xsl:when>
            <xsl:when test="$var = 'SRS1SDEL'">Delay % ANY</xsl:when>
            <xsl:when test="$var = 'SRS2SDEL'">Delay % COMM</xsl:when>
            <xsl:when test="$var = 'SRS3SDEL'">Delay % LOCL</xsl:when>
            <xsl:when test="$var = 'SRS4SDEL'">Delay % VIO</xsl:when>
            <xsl:when test="$var = 'SRS5SDEL'">Delay % SWAP</xsl:when>
            <xsl:when test="$var = 'SRS6SDEL'">Delay % OUTR</xsl:when>
            <xsl:when test="$var = 'SRS7SDEL'">Delay % XMEM</xsl:when>
            <xsl:when test="$var = 'SRS8SDEL'">Delay % HIPR</xsl:when>
            <xsl:when test="$var = 'SRS9SDEL'">Delay % OTHER</xsl:when>
            <xsl:when test="$var = 'SRSPACTV'">ACTV Frames</xsl:when>
            <xsl:when test="$var = 'SRSPIDLE'">IDLE Frames</xsl:when>
            <xsl:when test="$var = 'SRSPFIXD'">FIXED Frames</xsl:when>
            <xsl:when test="$var = 'SRSPPGIN'">Page-In Rate</xsl:when>
            <xsl:when test="$var = 'SRSDDSIN'">DDS Group</xsl:when>
            <xsl:when test="$var = 'SRSDDSIT'">DDS Type</xsl:when>
            <xsl:when test="$var = 'SRSDDSIP'">DDS Period</xsl:when>
            <!-- STORR -->
            <xsl:when test="$var = 'SRRVOLVC'">Volume Serial</xsl:when>
            <xsl:when test="$var = 'SRRDEVTY'">Device Type</xsl:when>
            <xsl:when test="$var = 'SRRCUTY'">Control Unit Type</xsl:when>
            <xsl:when test="$var = 'SRREXPCT'">Number of Exposures</xsl:when>
            <xsl:when test="$var = 'SRRUSVC'">% Using</xsl:when>
            <xsl:when test="$var = 'SRRA1VC'">% Active</xsl:when>
            <xsl:when test="$var = 'SRRA2VC'">% Connect</xsl:when>
            <xsl:when test="$var = 'SRRA3VC'">% Disconnect</xsl:when>
            <xsl:when test="$var = 'SRRA4VC'">% Pending</xsl:when>
            <xsl:when test="$var = 'SRRPDLYR'">Delay Reasons</xsl:when>
            <xsl:when test="$var = 'SRRPDLYP'">% Delay</xsl:when>
            <xsl:when test="$var = 'SRRA5VC'">% Delay DB</xsl:when>
            <xsl:when test="$var = 'SRRA8VC'">% Delay CMR</xsl:when>
            <xsl:when test="$var = 'SRRSPTVC'">Space Type</xsl:when>
            <xsl:when test="$var = 'SRRAUTOT'">Active Users Total</xsl:when>
            <xsl:when test="$var = 'SRRAULOC'">Active Users Local</xsl:when>
            <xsl:when test="$var = 'SRRAUSWP'">Active Users Swap</xsl:when>
            <xsl:when test="$var = 'SRRAUCOM'">Active Users Comm</xsl:when>
            <!-- ENQ -->
            <xsl:when test="$var = 'ENQPWJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'ENQPODEL'">% Delay</xsl:when>
            <xsl:when test="$var = 'ENQPRDEL'">% Resouce Waiting</xsl:when>
            <xsl:when test="$var = 'ENQPWSTT'">Resource Waiting Status</xsl:when>
            <xsl:when test="$var = 'ENQPMAJS'">Major/Minor Names (Scope)</xsl:when>
            <xsl:when test="$var = 'ENQPHDEL'">% Holding</xsl:when>
            <xsl:when test="$var = 'ENQPHJOB'">Holding Job/System Name</xsl:when>
            <xsl:when test="$var = 'ENQPHSTT'">Holding Job Status</xsl:when>
            <xsl:when test="$var = 'ENQPWJID'">JES ID</xsl:when>
            <!-- HSM and JES -->
            <xsl:when test="$var = 'HJSPJOB'">Job Name</xsl:when>
            <xsl:when test="$var = 'HJSPODEL'">% Total Delay</xsl:when>
            <xsl:when test="$var = 'HJS1FDEL'">% Delay (1)</xsl:when>
            <xsl:when test="$var = 'HJS1FCNR'">Function Code (1)</xsl:when>
            <xsl:when test="$var = 'HJS1EXPL'">Explanation (1)</xsl:when>
            <xsl:when test="$var = 'HJS2FDEL'">% Delay (2)</xsl:when>
            <xsl:when test="$var = 'HJS2FCNR'">Function Code (2)</xsl:when>
            <xsl:when test="$var = 'HJS2EXPL'">Explanation (2)</xsl:when>
            <xsl:when test="$var = 'HJSPJID'">JES ID</xsl:when>
            <!-- CACHSUM -->
            <xsl:when test="$var = 'CASHCDAT'">CDate</xsl:when>
            <xsl:when test="$var = 'CASHCTIM'">CTime</xsl:when>
            <xsl:when test="$var = 'CASHCRNG'">CRange</xsl:when>
            <xsl:when test="$var = 'CASPSSID'">Subsystem ID</xsl:when>
            <xsl:when test="$var = 'CASPCUID'">CU Number</xsl:when>
            <xsl:when test="$var = 'CASPTYPM'">Type&#160;Model</xsl:when>
            <xsl:when test="$var = 'CASPSIZE'">Subsystem Storage Size</xsl:when>
            <xsl:when test="$var = 'CASPIO'">I/O Rate</xsl:when>
            <xsl:when test="$var = 'CASPHITP'">Hit %</xsl:when>
            <xsl:when test="$var = 'CASPHIT'">Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASPMTOT'">Total Miss Rate</xsl:when>
            <xsl:when test="$var = 'CASPMSTG'">Stage Miss Rate</xsl:when>
            <xsl:when test="$var = 'CASPREAP'">Read %</xsl:when>
            <xsl:when test="$var = 'CASPSEQ'">Sequential Rate</xsl:when>
            <xsl:when test="$var = 'CASPASYN'">Async Rate</xsl:when>
            <xsl:when test="$var = 'CASPOFF'">Off Rate</xsl:when>
            <xsl:when test="$var = 'CASNRRA'">Normal Read Rate</xsl:when>
            <xsl:when test="$var = 'CASNRHI'">Normal Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASNRHIP'">Normal Read Hit %</xsl:when>
            <xsl:when test="$var = 'CASNWRA'">Normal Write Rate</xsl:when>
            <xsl:when test="$var = 'CASNWFA'">Normal Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CASNWHI'">Normal Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASNWHIP'">Normal Write Hit %</xsl:when>
            <xsl:when test="$var = 'CASNREAP'">Normal Read %</xsl:when>
            <xsl:when test="$var = 'CASNTRA'">Normal Tracks</xsl:when>
            <xsl:when test="$var = 'CASSRRA'">Sequential Read Rate</xsl:when>
            <xsl:when test="$var = 'CASSRHI'">Sequential Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASSRHIP'">Sequential Read Hit %</xsl:when>
            <xsl:when test="$var = 'CASSWRA'">Sequential Write Rate</xsl:when>
            <xsl:when test="$var = 'CASSWFA'">Sequential Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CASSWHI'">Sequential Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASSWHIP'">Sequential Write Hit %</xsl:when>
            <xsl:when test="$var = 'CASSREAP'">Sequential Read %</xsl:when>
            <xsl:when test="$var = 'CASSTRA'">Sequential Tracks</xsl:when>
            <xsl:when test="$var = 'CASCRRA'">CFW Read Rate</xsl:when>
            <xsl:when test="$var = 'CASCRHI'">CFW Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASCRHIP'">CFW Read Hit %</xsl:when>
            <xsl:when test="$var = 'CASCWRA'">CFW Write Rate</xsl:when>
            <xsl:when test="$var = 'CASCWFA'">CFW Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CASCWHI'">CFW Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASCWHIP'">CFW Write Hit %</xsl:when>
            <xsl:when test="$var = 'CASCREAP'">CFW Read %</xsl:when>
            <xsl:when test="$var = 'CASTRRA'">Total Read Rate</xsl:when>
            <xsl:when test="$var = 'CASTRHI'">Total Read Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASTRHIP'">Total Read Hit %</xsl:when>
            <xsl:when test="$var = 'CASTWRA'">Total Write Rate</xsl:when>
            <xsl:when test="$var = 'CASTWFA'">Total Write Fast Rate</xsl:when>
            <xsl:when test="$var = 'CASTWHI'">Total Write Hit Rate</xsl:when>
            <xsl:when test="$var = 'CASTWHIP'">Total Write Hit %</xsl:when>
            <xsl:when test="$var = 'CASTREAP'">Total Read %</xsl:when>
            <xsl:when test="$var = 'CASMCACH'">Caching Status</xsl:when>
            <xsl:when test="$var = 'CASMCCON'">Cache Configured</xsl:when>
            <xsl:when test="$var = 'CASMCAVL'">Cache Available</xsl:when>
            <xsl:when test="$var = 'CASMCOFF'">Cache Offline</xsl:when>
            <xsl:when test="$var = 'CASMCPIN'">Cache Pinned</xsl:when>
            <xsl:when test="$var = 'CASMNVS'">NVS Status</xsl:when>
            <xsl:when test="$var = 'CASMNCON'">NVS Configured</xsl:when>
            <xsl:when test="$var = 'CASMNPIN'">NVS Pinned</xsl:when>
            <xsl:otherwise>
                <xsl:value-of select="$var"/>
            </xsl:otherwise>
        </xsl:choose>
    </xsl:template>
</xsl:stylesheet>
