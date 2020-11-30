// Copyright (c) 2020 Rakuten, Inc. All rights reserved.
// Licensed under the MIt License.
// License that can be found in the LICENSE file.

package passenger

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// パース用ダミーXMLファイル
	TestXML = `
<?xml version="1.0" encoding="iso-8859-1"?>
<info version="3">
   <passenger_version>6.0.5</passenger_version>
   <group_count>1</group_count>
   <process_count>6</process_count>
   <max>9</max>
   <capacity_used>6</capacity_used>
   <get_wait_list_size>0</get_wait_list_size>
   <supergroups>
      <supergroup>
         <name>/git/ruby/passenger-ruby-rails (development)</name>
         <state>READY</state>
         <get_wait_list_size>0</get_wait_list_size>
         <capacity_used>6</capacity_used>
         <group default="true">
            <name>/git/ruby/passenger-ruby-rails (development)</name>
            <component_name>/git/ruby/passenger-ruby-rails (development)</component_name>
            <app_root>/git/ruby/passenger-ruby-rails</app_root>
            <app_type>ruby</app_type>
            <environment>development</environment>
            <uuid>Vr5AkjIVCo8h6G43cZBB</uuid>
            <enabled_process_count>6</enabled_process_count>
            <disabling_process_count>0</disabling_process_count>
            <disabled_process_count>0</disabled_process_count>
            <capacity_used>6</capacity_used>
            <get_wait_list_size>94</get_wait_list_size>
            <disable_wait_list_size>0</disable_wait_list_size>
            <processes_being_spawned>0</processes_being_spawned>
            <life_status>ALIVE</life_status>
            <user>user.name</user>
            <uid>2077423437</uid>
            <group>User Group</group>
            <gid>679754705</gid>
            <options>
               <app_root>/git/ruby/passenger-ruby-rails</app_root>
               <app_group_name>/git/ruby/passenger-ruby-rails (development)</app_group_name>
               <app_type>ruby</app_type>
               <start_command>ruby /Users/user.name/.rbenv/versions/2.7.0/lib/ruby/gems/2.7.0/gems/passenger-6.0.5/src/helper-scripts/rack-loader.rb</start_command>
               <startup_file>/git/ruby/passenger-ruby-rails/config.ru</startup_file>
               <log_level>3</log_level>
               <start_timeout>90000</start_timeout>
               <environment>development</environment>
               <base_uri>/</base_uri>
               <spawn_method>smart</spawn_method>
               <default_user>nobody</default_user>
               <default_group>nobody</default_group>
               <restart_dir>tmp</restart_dir>
               <integration_mode>standalone</integration_mode>
               <ruby>ruby</ruby>
               <python>python</python>
               <nodejs>node</nodejs>
               <debugger>false</debugger>
               <min_processes>1</min_processes>
               <max_processes>0</max_processes>
               <max_preloader_idle_time>300</max_preloader_idle_time>
               <max_out_of_band_work_instances>1</max_out_of_band_work_instances>
               <sticky_sessions_cookie_attributes>SameSite=Lax; Secure;</sticky_sessions_cookie_attributes>
            </options>
            <processes>
               <process>
                  <pid>4530</pid>
                  <sticky_session_id>931874451</sticky_session_id>
                  <gupid>xxxxx</gupid>
                  <concurrency>1</concurrency>
                  <sessions>1</sessions>
                  <busyness>2147483647</busyness>
                  <processed>498</processed>
                  <spawner_creation_time>1593644332662768</spawner_creation_time>
                  <spawn_start_time>1593644334798915</spawn_start_time>
                  <spawn_end_time>1593644334984986</spawn_end_time>
                  <last_used>1593646898924636</last_used>
                  <last_used_desc>0s ago</last_used_desc>
                  <uptime>42m 44s</uptime>
                  <life_status>ALIVE</life_status>
                  <enabled>ENABLED</enabled>
                  <has_metrics>true</has_metrics>
                  <cpu>0</cpu>
                  <rss>59444</rss>
                  <pss>-1</pss>
                  <private_dirty>-1</private_dirty>
                  <swap>-1</swap>
                  <real_memory>59444</real_memory>
                  <vmsize>4476308</vmsize>
                  <process_group_id>4481</process_group_id>
                  <command>Passenger AppPreloader: /git/ruby/passenger-ruby-rails (forking...)</command>
               </process>
               <process>
                  <pid>8222</pid>
               </process>
            </processes>
         </group>
      </supergroup>
   </supergroups>
</info>`
)

func TestParsePoolInfoSuccess(t *testing.T) {
	r := strings.NewReader(TestXML)

	info, err := ParsePoolInfo(r)
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "6.0.5", info.PassengerVersion)
	assert.Equal(t, 1, info.GroupCount)
	assert.Equal(t, 6, info.ProcessCount)
	assert.Equal(t, 9, info.Max)
	assert.Equal(t, 1, len(info.SuperGroups))

	sg := info.SuperGroups[0]
	assert.Equal(t, "/git/ruby/passenger-ruby-rails (development)", sg.Name)
	assert.Equal(t, "READY", sg.State)
	assert.Equal(t, 0, sg.GetWaitListSize)
	assert.Equal(t, 6, sg.CapacityUsed)

	g := sg.Group
	assert.Equal(t, "/git/ruby/passenger-ruby-rails (development)", g.Name)
	assert.Equal(t, "Vr5AkjIVCo8h6G43cZBB", g.UUID)
	assert.Equal(t, "ALIVE", g.LifeStatus)
	assert.Equal(t, 94, g.GetWaitListSize)
	assert.Equal(t, 2, len(g.Processes))

	p := g.Processes[0]
	assert.Equal(t, "4530", p.PID)
	assert.Equal(t, "ENABLED", p.Enabled)
	assert.Equal(t, "ALIVE", p.LifeStatus)
	assert.Equal(t, float64(498), p.Processed)
	assert.Equal(t, int64(59444), p.RSS)
	assert.Equal(t, int64(59444), p.RealMemory)
	assert.Equal(t, int64(4476308), p.VMSize)

	p = g.Processes[1]
	assert.Equal(t, "8222", p.PID)
	assert.Equal(t, "", p.Enabled)
	assert.Equal(t, "", p.LifeStatus)
	assert.Equal(t, float64(0), p.Processed)
	assert.Equal(t, int64(0), p.RSS)
	assert.Equal(t, int64(0), p.RealMemory)
	assert.Equal(t, int64(0), p.VMSize)
}

func TestParsePoolInfoParseFailure(t *testing.T) {
	contents := []string{"NOT FOUND", ""}
	for _, content := range contents {
		r := strings.NewReader(content)
		info, err := ParsePoolInfo(r)
		assert.Nil(t, info)
		assert.NotNil(t, err)
	}
}
