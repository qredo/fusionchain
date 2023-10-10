local default = import 'default.jsonnet';

default {
  'qredofusiontestnet-0'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
