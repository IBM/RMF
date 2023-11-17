// Generated from RMFQuery.g4 by ANTLR 4.8
// jshint ignore: start
var antlr4 = require('antlr4/index');

// This class defines a complete generic visitor for a parse tree produced by RMFQueryParser.

function RMFQueryVisitor() {
	antlr4.tree.ParseTreeVisitor.call(this);
	return this;
}

RMFQueryVisitor.prototype = Object.create(antlr4.tree.ParseTreeVisitor.prototype);
RMFQueryVisitor.prototype.constructor = RMFQueryVisitor;

// Visit a parse tree produced by RMFQueryParser#rmfquery.
RMFQueryVisitor.prototype.visitRmfquery = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#rmfqueryReport.
RMFQueryVisitor.prototype.visitRmfqueryReport = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#rmfqueryReportName.
RMFQueryVisitor.prototype.visitRmfqueryReportName = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#rmfqueryMetrices.
RMFQueryVisitor.prototype.visitRmfqueryMetrices = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#rmfqueryWithQul.
RMFQueryVisitor.prototype.visitRmfqueryWithQul = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#rmfqueryWoutQul.
RMFQueryVisitor.prototype.visitRmfqueryWoutQul = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#qualifications.
RMFQueryVisitor.prototype.visitQualifications = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#morequalification.
RMFQueryVisitor.prototype.visitMorequalification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#qualification.
RMFQueryVisitor.prototype.visitQualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#ulq_Qualification.
RMFQueryVisitor.prototype.visitUlq_Qualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#name_Qualification.
RMFQueryVisitor.prototype.visitName_Qualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#name_options.
RMFQueryVisitor.prototype.visitName_options = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#only_string.
RMFQueryVisitor.prototype.visitOnly_string = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#string_withdot.
RMFQueryVisitor.prototype.visitString_withdot = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#string_withdot_withnumber.
RMFQueryVisitor.prototype.visitString_withdot_withnumber = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#dot_withstring.
RMFQueryVisitor.prototype.visitDot_withstring = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#resource_withstring.
RMFQueryVisitor.prototype.visitResource_withstring = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#range_Qualification.
RMFQueryVisitor.prototype.visitRange_Qualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#range_Value.
RMFQueryVisitor.prototype.visitRange_Value = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#filter_Qualification.
RMFQueryVisitor.prototype.visitFilter_Qualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#filter_options.
RMFQueryVisitor.prototype.visitFilter_options = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#pat_filter.
RMFQueryVisitor.prototype.visitPat_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#lb_filter.
RMFQueryVisitor.prototype.visitLb_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#ub_filter.
RMFQueryVisitor.prototype.visitUb_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#hi_filter.
RMFQueryVisitor.prototype.visitHi_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#lo_filter.
RMFQueryVisitor.prototype.visitLo_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#ord_filter.
RMFQueryVisitor.prototype.visitOrd_filter = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#workscope_Qualification.
RMFQueryVisitor.prototype.visitWorkscope_Qualification = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#workscope_options.
RMFQueryVisitor.prototype.visitWorkscope_options = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#workscope_ulq.
RMFQueryVisitor.prototype.visitWorkscope_ulq = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#workscope_name.
RMFQueryVisitor.prototype.visitWorkscope_name = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#workscope_type.
RMFQueryVisitor.prototype.visitWorkscope_type = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#reptypename.
RMFQueryVisitor.prototype.visitReptypename = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#reptype.
RMFQueryVisitor.prototype.visitReptype = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#repname.
RMFQueryVisitor.prototype.visitRepname = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#restype.
RMFQueryVisitor.prototype.visitRestype = function(ctx) {
  return this.visitChildren(ctx);
};


// Visit a parse tree produced by RMFQueryParser#metrics.
RMFQueryVisitor.prototype.visitMetrics = function(ctx) {
  return this.visitChildren(ctx);
};



exports.RMFQueryVisitor = RMFQueryVisitor;
