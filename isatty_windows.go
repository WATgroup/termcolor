// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

//go:build windows

package color

func isTerminal(fd int) bool {

	if fd < 0 {
		return false
	}
	return true
}
