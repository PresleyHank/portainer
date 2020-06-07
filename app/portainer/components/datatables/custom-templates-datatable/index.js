import angular from 'angular';
import CustomTemplatesDatatableController from './customTemplatesDatatableController';

angular.module('portainer.app').component('customTemplatesDatatable', {
  templateUrl: './customTemplatesDatatable.html',
  controller: CustomTemplatesDatatableController,
  bindings: {
    isAdmin: '<',
    currentUserId: '<',
    titleText: '@',
    titleIcon: '@',
    dataset: '<',
    tableKey: '@',
    orderBy: '@',
    reverseOrder: '<',
    accessManagement: '<',
    removeAction: '<',
  },
});
