CERNBox DB schema
=================

The CERNBox database schema inherits the ownCloud 10 schema and adds Reva-specific tables.

This file describes the state of the database following a reverse engineering analysis.

The following tables with their corresponding cardinalities are used in production and referenced in Reva as of July 2023:

* oc_preferences            # 65K
* oc_share                  # 490K
* oc_share_status           # 24K

* cernbox_project_mapping   # 1K
* cbox_metadata             # 10K, to store favourites
* cbox_otg_ocis             # empty, to be deprecated in favour of notifications

* notifications
* notification_recipients

* ocm_access_method_webapp
* ocm_access_method_webdav
* ocm_protocol_transfer
* ocm_protocol_webapp
* ocm_protocol_webdav
* ocm_received_share_protocols
* ocm_received_shares
* ocm_remote_users
* ocm_shares
* ocm_shares_access_methods
* ocm_tokens

The following tables contain several records (possibly used by ownCloud in the past) but are not used by Reva:

* oc_mounts
* oc_share_acl
* oc_share_acl_old
* oc_addressbooks
* oc_accounts

The following tables were used in the past but not any longer:

* cbox_canary
* cbox_office_engine

The following table was populated by hand and it is not used by Reva:

* oc_sciencemesh

