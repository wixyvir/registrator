# Registrator

Service registry bridge for Docker, sponsored by [Weave](http://weave.works).

Registrator automatically registers and deregisters services for any Docker
container by inspecting containers as they come online. Registrator
supports pluggable service registries, which currently includes
[Consul](http://www.consul.io/), [etcd](https://github.com/coreos/etcd) and
[SkyDNS 2](https://github.com/skynetservices/skydns/).

Full documentation available at http://gliderlabs.com/registrator

## This gliderlabs/registrator fork

**This fork TLS support [has been merged](https://github.com/gliderlabs/registrator/pull/394) in the official gliderlabs/registrator project**

This fork add TLS support for Consul Backend.

Run:
```
docker run -d --name=registrator --net=host  -v /var/run/docker.sock:/tmp/docker.sock -v /etc/certificates/:/etc/certificates/ -e CONSUL_CACERT=/etc/pmsipilot/docker-ca.crt -e CONSUL_TLSCERT=/etc/certificates/consul-client.crt -e CONSUL_TLSKEY=/etc/certificates/consul-client.key cyprien/registrator consul-tls://srv-consul.local:8543
```

## Getting Registrator

Get the latest release, master, or any version of Registrator via [Docker Hub](https://registry.hub.docker.com/u/gliderlabs/registrator/):

	$ docker pull gliderlabs/registrator:latest

Latest tag always points to the latest release. There is also a `:master` tag
and version tags to pin to specific releases.

## Using Registrator

The quickest way to see Registrator in action is our
[Quickstart](https://gliderlabs.com/registrator/latest/user/quickstart)
tutorial. Otherwise, jump to the [Run
Reference](https://gliderlabs.com/registrator/latest/user/run) in the User
Guide. Typically, running Registrator looks like this:

    $ docker run -d \
        --name=registrator \
        --net=host \
        --volume=/var/run/docker.sock:/tmp/docker.sock \
        gliderlabs/registrator:latest \
          consul://localhost:8500

## Contributing

Pull requests are welcome! We recommend getting feedback before starting by
opening a [GitHub issue](https://github.com/gliderlabs/registrator/issues) or
discussing in [Slack](http://glider-slackin.herokuapp.com/).

Also check out our Developer Guide on [Contributing
Backends](https://gliderlabs.com/registrator/latest/dev/backends) and [Staging
Releases](https://gliderlabs.com/registrator/latest/dev/releases).

## Sponsors and Thanks

Ongoing support of this project is made possible by [Weave](http://weave.works), the easiest way to connect, observe and control your containers. Big thanks to Michael Crosby for
[skydock](https://github.com/crosbymichael/skydock) and the Consul mailing list
for inspiration.

For a full list of sponsors, see
[SPONSORS](https://github.com/gliderlabs/registrator/blob/master/SPONSORS).

## License

MIT

<img src="https://ga-beacon.appspot.com/UA-58928488-2/registrator/readme?pixel" />
