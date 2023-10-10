local config = import 'default.jsonnet';

config {
  'qredofusiontestnet-0'+: {
    validators: super.validators[0:1] + [{
      name: 'fullnode',
    }],
  },
}
